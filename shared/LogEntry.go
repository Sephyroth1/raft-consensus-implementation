package shared

type LogEntry struct {
	LogIndex int
	Term     int
	Command  string
}
