/*
	Package rpc is used to store the structure of the RPC's

# AppendEntires is the heartbeat and the Log Entry RPC

# It Includes the Arguments and Results of the RPC's

Arguments:
  - term - leader's term
  - leaderId - so followers can redirect clients
  - prevLogIndex - index of log entry immediately preceding new ones
  - prevLogTerm - term of log entry immediately preceding new ones
  - entries[] - log entries to store (empty for heartbeats;
  - may send more than one for efficiency)

Results:
  - term - currentTerm; for leader to update itself
  - success - true if follower contained entry matching prevLogIndex and prevLogTerm
*/
package rpc

import (
	netrpc "net/rpc"

	"consensus.com/shared"
)

type AppendEntries struct {
	Term         int
	LeaderID     int
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []shared.LogEntry
	LeaderCommit int
}

type AppendEntriesResult struct {
	Term    int
	Success bool
}

func sendAppendEntries(client *netrpc.Client, args *AppendEntries, result *AppendEntriesResult) error {
	err := client.Call("Raft.AppendEntries", args, result)
	if err != nil {
		return err
	}
	return nil
}
