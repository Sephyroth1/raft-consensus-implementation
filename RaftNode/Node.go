/*
	  Package raftnode is used to store the structure of the Raft Node
			  A Raft Node is a server that replicates the log and executes commands
				It has a persistent state, a volatile state and a volatile state on leaders

				Persistent State:
				- currentTerm - latest term server has seen (initialized to 0)
				- votedFor - candidate that received vote in current term (null if none)
				- log[] - log entries (entries with smaller indexes are at head of array)

				Volatile State:
				- commitIndex - index of highest log entry known to be committed
				- lastApplied - index of highest log entry applied to state machine

			  Volatile State on Leaders:
				- nextIndex[] - for each server, index of next log entry to send to that server
								(initialized to leader last log index + 1)
				- matchIndex[] - for each server, index of highest log entry known to be
								replicated on server (initialized to 0)
*/
package raftnode

import (
	"slices"
	"sync"
	"time"

	netrpc "net/rpc"

	config "consensus.com/Config"
	rpc "consensus.com/RPC"
	"consensus.com/shared"
)

type Node struct {
	Mu         *sync.Mutex
	Node       shared.NodeEntry
	Election   *time.Timer
	Heartbeat  *time.Timer
	MinTimeout int
	MaxTimeout int
	RPCChan    chan shared.RPCEvent
	Votes      int
}

func NewNode(node shared.NodeEntry, con config.ConfigEnv) *shared.NodeEntry {
	return &shared.NodeEntry{
		ID:          node.ID,
		Address:     node.Address,
		Port:        node.Port,
		Role:        node.Role,
		CurrentTerm: node.CurrentTerm,
		VotedFor:    node.VotedFor,
		Log:         node.Log,
		CommitIndex: node.CommitIndex,
		LastApplied: node.LastApplied,
		NextIndex:   node.NextIndex,
		MatchIndex:  node.MatchIndex,
	}
}

func ConvertNode(node shared.NodeEntry) Node {
	config := config.ConfigEnv{}
	return Node{
		Mu:       &sync.Mutex{},
		Node:     node,
		Election: time.NewTimer(config.RandomizedTimeout()),
	}
}

func (r *Node) AppendEntries(args *rpc.AppendEntries, reply *rpc.AppendEntriesResult) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	if args.Term < r.Node.CurrentTerm {
		reply.Term = r.Node.CurrentTerm
		reply.Success = false
		return
	}

	if args.PrevLogIndex > 0 {
		if args.PrevLogIndex == len(r.Node.Log) && r.Node.Log[args.PrevLogIndex].Term != args.PrevLogTerm {
			reply.Success = false
			r.Node.Log = slices.Delete(r.Node.Log, args.PrevLogIndex, len(r.Node.Log))
			return
		}
	} else {
		if (args.PrevLogIndex < 0 || args.PrevLogIndex >= len(r.Node.Log)) || r.Node.Log[args.PrevLogIndex].Term != args.PrevLogTerm {
			reply.Success = false
			return
		}
	}
	r.Node.Log = append(r.Node.Log, args.Entries...)
	if args.LeaderCommit > r.Node.CommitIndex {
		r.Node.CommitIndex = min(args.LeaderCommit, len(r.Node.Log)-1)
	}
	r.RPCChan <- shared.RPCEvent{
		Type: "AppendEntries",
		Data: args,
	}
	reply.Term = r.Node.CurrentTerm
	reply.Success = true
}

func (r *Node) RequestVote(args *rpc.RequestVote, reply *rpc.RequestVoteResult) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	if args.Term < r.Node.CurrentTerm {
		reply.Term = r.Node.CurrentTerm
		reply.VoteGranted = false
		return
	}

	if r.Node.VotedFor == 0 || args.CandidateID == 0 {
		reply.Term = r.Node.CurrentTerm
		reply.VoteGranted = false
		return
	}

	r.RPCChan <- shared.RPCEvent{
		Type: "RequestVote",
		Data: args,
	}
	reply.Term = r.Node.CurrentTerm
	reply.VoteGranted = true
}

func (r *Node) Run() {
	for {
		select {
		case <-r.Election.C:
			r.Node.CurrentTerm += 1
			r.Node.Role = shared.Candidate
			args := rpc.RequestVote{
				r.Node.CurrentTerm,
				r.Node.ID,
				r.Node.Log[len(r.Node.Log)-1].LogIndex,
				r.Node.Log[len(r.Node.Log)-1].Term,
			}
			reply := rpc.RequestVoteResult{}
			r.RequestVote(&args, &reply)
		}
	}
}

func (r *Node) HandleVoteResponse(
	client *netrpc.Client,
	peer int,
	electionTerm int,
	reply *rpc.RequestVoteResult,
) {
	con := config.ConfigEnv{}
	if reply.Term > electionTerm {
		r.Node.CurrentTerm = reply.Term
		r.Node.Role = shared.Follower
		r.Election = time.NewTimer(con.RandomizedTimeout())
		return
	}

	if r.Node.Role != shared.Candidate {
		return
	}

	if electionTerm != r.Node.CurrentTerm {
		return
	}

	if reply.VoteGranted {
		r.Votes++
	}

	if r.Votes >= con.Majority {
		for ind, peer := range con.Peers {
			client.Go("Raft.Heartbeat", nil, nil)
			r.Node.NextIndex[ind] = len(r.Node.Log)
			r.Node.MatchIndex[ind] = 0
		}
		r.Node.Role = shared.Leader
	}
}
