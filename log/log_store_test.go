package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStoreLogs(t *testing.T) {
	logStore := NewLogMemStore()
	logs := []*Log{
		{
			Index:      uint64(1),
			Term:       uint64(1),
			Data:       []byte("hello"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(2),
			Term:       uint64(1),
			Data:       []byte("world"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(3),
			Term:       uint64(1),
			Data:       []byte("foo"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(4),
			Data:       []byte("bar"),
			AppendedAt: time.Now(),
		},
	}
	err := logStore.StoreLogs(logs)
	assert.NoError(t, err)

	for _, expected := range logs {
		result, err := logStore.GetLog(expected.Index)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	}
}

func TestStoreLog(t *testing.T) {
	logStore := NewLogMemStore()
	log := &Log{
		Index:      uint64(1),
		Term:       uint64(1),
		Data:       []byte("hello"),
		AppendedAt: time.Now(),
	}
	err := logStore.StoreLog(log)
	assert.NoError(t, err)

	result, err := logStore.GetLog(log.Index)
	assert.NoError(t, err)
	assert.Equal(t, log, result)
}

func TestGetLog_NonExistentIndex(t *testing.T) {
	logStore := NewLogMemStore()
	_, err := logStore.GetLog(uint64(1))
	assert.Error(t, err, "not found")
}

func TestLastIndex_ReturnZeroWhenNoLogs(t *testing.T) {
	logStore := NewLogMemStore()
	index, err := logStore.LastIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), index)
}

func TestLastIndex_ReturnsIndexOfLastAddedElement(t *testing.T) {
	logStore := NewLogMemStore()
	logs := []*Log{
		{
			Index:      uint64(1),
			Term:       uint64(1),
			Data:       []byte("hello"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(2),
			Term:       uint64(1),
			Data:       []byte("world"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(3),
			Term:       uint64(1),
			Data:       []byte("foo"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(4),
			Data:       []byte("bar"),
			AppendedAt: time.Now(),
		},
	}
	err := logStore.StoreLogs(logs)
	assert.NoError(t, err)

	index, err := logStore.LastIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(4), index)
}

func TestFirstIndex_ReturnZeroWhenNoLogs(t *testing.T) {
	logStore := NewLogMemStore()
	index, err := logStore.FirstIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), index)
}

func TestFirstIndex_ReturnsIndexOfFirstAddedElement(t *testing.T) {
	logStore := NewLogMemStore()
	logs := []*Log{
		{
			Index:      uint64(1),
			Term:       uint64(1),
			Data:       []byte("hello"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(2),
			Term:       uint64(1),
			Data:       []byte("world"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(3),
			Term:       uint64(1),
			Data:       []byte("foo"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(4),
			Data:       []byte("bar"),
			AppendedAt: time.Now(),
		},
	}
	err := logStore.StoreLogs(logs)
	assert.NoError(t, err)

	index, err := logStore.FirstIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), index)
}

func TestFirstIndex_ReturnsIndexOfTheFirstIndexNonSortedList(t *testing.T) {
	logStore := NewLogMemStore()
	logs := []*Log{

		{
			Index:      uint64(2),
			Term:       uint64(1),
			Data:       []byte("world"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(3),
			Term:       uint64(1),
			Data:       []byte("foo"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(1),
			Term:       uint64(1),
			Data:       []byte("hello"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(4),
			Data:       []byte("bar"),
			AppendedAt: time.Now(),
		},
	}
	err := logStore.StoreLogs(logs)
	assert.NoError(t, err)

	index, err := logStore.FirstIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), index)
}

func TestDeleteRange_RemovesRangeInclusive(t *testing.T) {
	logStore := NewLogMemStore()
	logs := []*Log{
		{
			Index:      uint64(2),
			Term:       uint64(1),
			Data:       []byte("world"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(3),
			Term:       uint64(1),
			Data:       []byte("foo"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(1),
			Term:       uint64(1),
			Data:       []byte("hello"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(4),
			Data:       []byte("bar"),
			AppendedAt: time.Now(),
		},
	}
	err := logStore.StoreLogs(logs)
	assert.NoError(t, err)

	err = logStore.DeleteRange(2, 3)
	assert.NoError(t, err)

	_, err = logStore.GetLog(1)
	assert.NoError(t, err)

	_, err = logStore.GetLog(2)
	assert.Error(t, err, "not found")

	_, err = logStore.GetLog(3)
	assert.Error(t, err, "not found")

	_, err = logStore.GetLog(4)
	assert.NoError(t, err)
}

func TestDeleteRange_UpdatesIndexesProperly(t *testing.T) {
	logStore := NewLogMemStore()
	logs := []*Log{
		{
			Index:      uint64(2),
			Term:       uint64(1),
			Data:       []byte("world"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(3),
			Term:       uint64(1),
			Data:       []byte("foo"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(1),
			Term:       uint64(1),
			Data:       []byte("hello"),
			AppendedAt: time.Now(),
		},
		{
			Index:      uint64(4),
			Data:       []byte("bar"),
			AppendedAt: time.Now(),
		},
	}
	err := logStore.StoreLogs(logs)
	assert.NoError(t, err)

	err = logStore.DeleteRange(2, 3)
	assert.NoError(t, err)

	firstIndex, err := logStore.FirstIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), firstIndex)

	lastIndex, err := logStore.LastIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(4), lastIndex)

	err = logStore.DeleteRange(4, 4)
	assert.NoError(t, err)

	firstIndex, err = logStore.FirstIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), firstIndex)

	lastIndex, err = logStore.LastIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), lastIndex)

	err = logStore.DeleteRange(1, 1)
	assert.NoError(t, err)

	firstIndex, err = logStore.FirstIndex()
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), firstIndex)

	lastIndex, err = logStore.LastIndex()
	assert.NoError(t, err)
}
