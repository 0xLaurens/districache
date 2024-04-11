package proto

import (
	"bytes"
	"encoding/binary"
)

type Command byte

const (
	CmdGet Command = iota
	CmdSet
	CmdDelete
)

type CommandSet struct {
	Key   []byte
	Value []byte
	TTL   int
}

func (c *CommandSet) Bytes() []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, CmdSet)

	_ = binary.Write(buf, binary.LittleEndian, uint32(len(c.Key)))
	_ = binary.Write(buf, binary.LittleEndian, c.Key)

	_ = binary.Write(buf, binary.LittleEndian, uint32(len(c.Value)))
	_ = binary.Write(buf, binary.LittleEndian, c.Value)

	_ = binary.Write(buf, binary.LittleEndian, c.TTL)

	return buf.Bytes()
}

type CommandGet struct {
	Key []byte
}

func (c *CommandGet) Bytes() []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, CmdGet)

	_ = binary.Write(buf, binary.LittleEndian, uint32(len(c.Key)))
	_ = binary.Write(buf, binary.LittleEndian, c.Key)

	return buf.Bytes()
}

type CommandDelete struct {
	Key []byte
}

func (c *CommandDelete) Bytes() []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, CmdDelete)

	_ = binary.Write(buf, binary.LittleEndian, uint32(len(c.Key)))
	_ = binary.Write(buf, binary.LittleEndian, c.Key)

	return buf.Bytes()
}
