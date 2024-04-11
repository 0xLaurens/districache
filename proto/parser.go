package proto

import (
	"encoding/binary"
	"fmt"
	"io"
)

func ParseCommand(r io.Reader) (any, error) {
	var cmd Command
	if err := binary.Read(r, binary.LittleEndian, &cmd); err != nil {
		return nil, err
	}

	switch cmd {
	case CmdGet:
		return parseGetCommand(r)
	case CmdDelete:
		return parseDeleteCommand(r)
	case CmdSet:
		return parseSetCommand(r)
	default:
		return nil, fmt.Errorf("unknown cmd")
	}
}

func parseGetCommand(r io.Reader) (*CommandGet, error) {
	cmd := &CommandGet{}

	var keyLen uint32
	_ = binary.Read(r, binary.LittleEndian, &keyLen)
	cmd.Key = make([]byte, keyLen)
	_ = binary.Read(r, binary.LittleEndian, &cmd.Key)

	return cmd, nil
}

func parseDeleteCommand(r io.Reader) (*CommandDelete, error) {
	cmd := &CommandDelete{}

	var keyLen uint32
	_ = binary.Read(r, binary.LittleEndian, &keyLen)

	cmd.Key = make([]byte, keyLen)
	_ = binary.Read(r, binary.LittleEndian, &cmd.Key)

	return cmd, nil
}

func parseSetCommand(r io.Reader) (*CommandSet, error) {
	cmd := &CommandSet{}

	var keyLen uint32
	_ = binary.Read(r, binary.LittleEndian, &keyLen)

	cmd.Key = make([]byte, keyLen)
	_ = binary.Read(r, binary.LittleEndian, &cmd.Key)

	var valLen uint32
	_ = binary.Read(r, binary.LittleEndian, &valLen)

	cmd.Value = make([]byte, valLen)
	_ = binary.Read(r, binary.LittleEndian, &cmd.Value)

	_ = binary.Read(r, binary.LittleEndian, &cmd.TTL)

	return cmd, nil
}
