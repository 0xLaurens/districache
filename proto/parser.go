package proto

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

func ParseCommand(r io.Reader) (any, error) {
	var cmd Command
	err := binary.Read(r, binary.LittleEndian, &cmd)
	if err != nil {
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

func parseGetCommand(r io.Reader) (CommandGet, error) {
	log.Println("parse get command")
	cmd := CommandGet{}

	var keyLen int32
	_ = binary.Read(r, binary.LittleEndian, &keyLen)

	cmd.Key = make([]byte, keyLen)
	_ = binary.Read(r, binary.LittleEndian, &cmd.Key)

	return cmd, nil
}

func parseDeleteCommand(r io.Reader) (CommandDelete, error) {
	log.Println("parse delete command")
	cmd := CommandDelete{}

	var keyLen int32
	_ = binary.Read(r, binary.LittleEndian, &keyLen)

	cmd.Key = make([]byte, keyLen)
	_ = binary.Read(r, binary.LittleEndian, &cmd.Key)

	return cmd, nil
}

func parseSetCommand(r io.Reader) (CommandSet, error) {
	log.Println("parse set command")
	cmd := CommandSet{}

	var keyLen int32
	_ = binary.Read(r, binary.LittleEndian, &keyLen)

	cmd.Key = make([]byte, keyLen)
	_ = binary.Read(r, binary.LittleEndian, &cmd.Key)

	var valLen int32
	_ = binary.Read(r, binary.LittleEndian, &valLen)

	cmd.Value = make([]byte, valLen)
	_ = binary.Read(r, binary.LittleEndian, &cmd.Value)

	_ = binary.Read(r, binary.LittleEndian, &cmd.TTL)

	return cmd, nil
}
