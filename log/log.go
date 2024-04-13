package log

import "time"

// Log is a data structure that holds the log entry
// of the distributed cache. It is used to keep track of
// the state of the cache and to replicate the state
// across the cluster.
type Log struct {
	// Index holds the index of the log entry
	Index uint64

	// Term holds the election term of the log entry
	Term uint64

	// Data holds the log entry's data
	Data []byte

	// AppendedAt stores the time a log was appended
	// important to reach a similar state in terms of ttl.
	// ex. if a log was appended 3 min ago and the ttl is 1 min the
	// action should not be included in the state
	AppendedAt time.Time
}

type LogStore interface {
	// FirstIndex returns the index of the first index written 0 if no logs
	FirstIndex() (uint64, error)

	// LastIndex returns the index of the last index written 0 if no logs
	LastIndex() (uint64, error)

	// StoreLog stores a single log entry
	StoreLog(log *Log) error

	// StoreLogs stores multiple log entries. Useful for removing a gap in the log entries.
	StoreLogs(logs []*Log) error

	// GetLog returns a log entry at the given index. returns a error if the log does not exist.
	GetLog(index uint64) (*Log, error)

	// DeleteRange deletes all log entries between start and end index. The range is inclusive.
	DeleteRange(min, max uint64) error
}
