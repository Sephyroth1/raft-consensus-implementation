# Raft Node State

- Raft Node States have 4 parts
- They are:

  - ## Persistent state of all nodes

  - ## Volatile state of all nodes

  - ## Volatile state on leaders

## Persistent State

- currentTerm - latest term server has seen (initialized to 0)
- votedFor - candidate that received vote in current term (null if none)
- log[] - log entries (entries with smaller indexes are at head of array)

## Volatile State

- commitIndex - index of highest log entry known to be committed
- lastApplied - index of highest log entry applied to state machine

## Volatile State on Leaders

- nextIndex[] - for each server, index of next log entry to send to that server
                (initialized to leader last log index + 1)
- matchIndex[] - for each server, index of highest log entry known to be
                replicated on server (initialized to 0)

Next Important to talk about is The RPC's and their structure

## RPC Structure

- For RPC to work, I need to design the RPC structure for each purpose
- First purpose is heartbeats and appendentries rpc

### AppendEntries RPC

#### Arguments

- term - leader's term
- leaderId - so followers can redirect clients
- prevLogIndex - index of log entry immediately preceding new ones
- prevLogTerm - term of log entry immediately preceding new ones

- entries[] - log entries to store (empty for heartbeats;
                                    may send more than one for efficiency)
- leaderCommit - leader's commitIndex

#### Results

- term - currentTerm; for leader to update itself
- success - true if follower contained entry matching prevLogIndex and prevLogTerm

### Receiver's Implementation

1. Reply False if term < currentTerm
2. Reply False if log doesn't contain an entry at prevLogIndex
whose term matches prevLogTerm
3. If an existing entry conflicts with a new one (same index but different terms),
delete the existing entry and all that follow it
4. Append any new entries not already in the log
5. If leaderCommit > commitIndex,
set commitIndex = min(leaderCommit, index of last new entry)
