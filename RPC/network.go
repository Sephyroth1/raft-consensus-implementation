/* Package rpc is the package that contains the networking code for the consensus algorithm.
* It produces a networking layer for the consensus algorithm to communicate with other nodes.
* It specifically uses TCP/IP for the networking layer.
* Also it uses RPC with the RPC structures to network
* */
package rpc

import (
	"net"
	netrpc "net/rpc"
	"strconv"

	"consensus.com/shared"
)

type Raft struct {
	RaftNode *shared.NodeEntry
}

func Register(r *Raft) {
	netrpc.Register(r)
}

func Listen(port int) (net.Listener, error) {
	raft := shared.NodeEntry{
		CurrentTerm: 0,
		VotedFor:    0,
		Log:         []shared.LogEntry{},
		CommitIndex: 0,
		LastApplied: 0,
		NextIndex:   []int{},
		MatchIndex:  []int{},
		Role:        shared.Follower,
	}
	netrpc.Register(new(raft))
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	go netrpc.Accept(listener)
	return listener, nil
}

func Dial(address string, port int) *netrpc.Client {
	client, err := netrpc.DialHTTP("tcp", address+":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	return client
}
