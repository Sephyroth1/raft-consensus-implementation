/*
		Package rpc is used to store the structure of the RPC's
	  RequestVote is an RPC structure that is used to ask for votes by Candidates

		Arguments:
			- term - candidate's term
			- candidateId - candidate's ID
			- lastLogIndex - index of candidate's last log entry
			- lastLogTerm - term of candidate's last log entry

		Results:
			- term - currentTerm; for candidate to update itself
			- voteGranted - true if candidate received vote
*/
package rpc

import (
	netrpc "net/rpc"
)

type RequestVote struct {
	Term         int
	CandidateID  int
	LastLogIndex int
	LastLogTerm  int
}

type RequestVoteResult struct {
	Term        int
	VoteGranted bool
}

func SendRequestVote(client *netrpc.Client, args *RequestVote, result *RequestVoteResult) error {
	err := client.Call("Raft.RequestVote", args, result)
	if err != nil {
		return err
	}
	return nil
}
