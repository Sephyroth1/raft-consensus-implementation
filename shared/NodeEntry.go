package shared

type NodeEntry struct {
	ID          int
	Address     string
	Port        int
	Role        Role
	CurrentTerm int
	VotedFor    int
	Log         []LogEntry
	CommitIndex int
	LastApplied int
	NextIndex   []int
	MatchIndex  []int
}
