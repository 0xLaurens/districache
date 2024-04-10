package proto

import (
	"bytes"
	"encoding/binary"
)

type Status byte

const (
	StatusOK Status = iota
	StatusError
	StatusNone
)

func (s Status) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusError:
		return "ERR"
	default:
		return "NONE"
	}
}

type ResponseSet struct {
	Status Status
}

func (r *ResponseSet) Bytes() []byte {
	buf := new(bytes.Buffer)

	_ = binary.Write(buf, binary.LittleEndian, r.Status)

	return buf.Bytes()
}

type ResponseDelete struct {
	Status Status
}

func (r *ResponseDelete) Bytes() []byte {
	buf := new(bytes.Buffer)

	_ = binary.Write(buf, binary.LittleEndian, r.Status)

	return buf.Bytes()
}

type ResponseGet struct {
	Status Status
	Value  []byte
}

func (r *ResponseGet) Bytes() []byte {
	buf := new(bytes.Buffer)

	_ = binary.Write(buf, binary.LittleEndian, r.Status)
	_ = binary.Write(buf, binary.LittleEndian, int32(len(r.Value)))
	_ = binary.Write(buf, binary.LittleEndian, r.Value)

	return buf.Bytes()
}
