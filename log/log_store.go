package log

import (
	"errors"
	"sync"
)

// LogMemStore is an in-memory implementation of the LogStore.
// should not be used like this in production, it is harder to restore a cluster to its original state.
type LogMemStore struct {
	firstIndex uint64
	lastIndex  uint64
	logs       map[uint64]*Log
	lock       sync.RWMutex
}

// NewLogMemStore returns an in-memory backend.
func NewLogMemStore() *LogMemStore {
	return &LogMemStore{
		firstIndex: 0,
		lastIndex:  0,
		logs:       make(map[uint64]*Log),
	}
}

func (l *LogMemStore) StoreLog(log *Log) error {
	return l.StoreLogs([]*Log{log})
}

func (l *LogMemStore) StoreLogs(logs []*Log) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	for _, log := range logs {
		l.logs[log.Index] = log
		l.lastIndex = log.Index

		if l.firstIndex == 0 {
			l.firstIndex = log.Index
		}

		if log.Index < l.firstIndex {
			l.firstIndex = log.Index
		}
	}

	return nil
}

func (l *LogMemStore) GetLog(index uint64) (*Log, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	log, ok := l.logs[index]
	if !ok {
		return nil, errors.New("not found")
	}

	return log, nil
}

func (l *LogMemStore) DeleteRange(min, max uint64) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	for i := min; i <= max; i++ {
		delete(l.logs, i)
	}
	if min <= l.firstIndex {
		l.firstIndex = max + 1
	}
	if max >= l.lastIndex {
		l.lastIndex = min - 1
	}
	if l.firstIndex > l.lastIndex {
		l.firstIndex = 0
		l.lastIndex = 0
	}

	return nil
}

func (l *LogMemStore) FirstIndex() (uint64, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.firstIndex, nil
}

func (l *LogMemStore) LastIndex() (uint64, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.lastIndex, nil
}
