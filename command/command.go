package command

import (
	"time"
)

type CMD string

const (
	CMDGet    CMD = "GET"
	CMDSet    CMD = "SET"
	CMDDelete CMD = "DEL"
)

type Command interface {
	GetType() CMD
	GetKey() []byte
}

type BaseCMD struct {
	Cmd CMD
	Key []byte
}

func (b BaseCMD) GetType() CMD {
	return b.Cmd
}

func (b BaseCMD) GetKey() []byte {
	return b.Key
}

type SetCMD struct {
	BaseCMD
	Value []byte
	TTL   time.Duration
}

func (g SetCMD) GetType() CMD {
	return g.Cmd
}

func (g SetCMD) GetKey() []byte {
	return g.Key
}
