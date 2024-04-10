package proto

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseGetCommand(t *testing.T) {
	cmd := &CommandGet{
		Key: []byte("hello"),
	}
	r := bytes.NewReader(cmd.Bytes())
	command, err := ParseCommand(r)
	assert.NoError(t, err)

	assert.Equal(t, cmd, command)
}

func TestParseSet(t *testing.T) {
	cmd := &CommandSet{
		Key:   []byte("hello"),
		Value: []byte("world"),
		TTL:   0,
	}
	r := bytes.NewReader(cmd.Bytes())
	command, err := ParseCommand(r)
	assert.NoError(t, err)

	assert.Equal(t, cmd, command)
}

func TestParseDelete(t *testing.T) {
	cmd := &CommandDelete{
		Key: []byte("hello"),
	}
	r := bytes.NewReader(cmd.Bytes())
	command, err := ParseCommand(r)
	assert.NoError(t, err)

	assert.Equal(t, cmd, command)
}

func BenchmarkParseCommand(b *testing.B) {
	cmd := &CommandSet{
		Key:   []byte("hello"),
		Value: []byte("world"),
		TTL:   0,
	}

	r := bytes.NewReader(cmd.Bytes())

	for i := 0; i < b.N; i++ {
		_, _ = ParseCommand(r)
	}
}
