package main

import (
	"encoding/binary"
	"github.com/0xlaurens/districache/proto"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return
	}
	SetCMD(conn, []byte("key"), []byte("value"), 0)
	resSet, _ := ParseSetResponse(conn)
	log.Println("", resSet)
	GetCMD(conn, []byte("key"))
	resGet, _ := ParseGetResponse(conn)
	log.Println(resGet)
	time.Sleep(300 * time.Millisecond)
}

func GetCMD(conn net.Conn, key []byte) {
	cmd := &proto.CommandGet{
		Key: key,
	}
	_, err := conn.Write(cmd.Bytes())
	if err != nil {
		return
	}
}

func SetCMD(conn net.Conn, key []byte, value []byte, TTL int) {
	cmd := &proto.CommandSet{
		Key:   key,
		Value: value,
		TTL:   TTL,
	}

	_, err := conn.Write(cmd.Bytes())
	if err != nil {
		return
	}
}

func DeleteCMD(conn net.Conn, key []byte) {
	cmd := &proto.CommandDelete{
		Key: key,
	}

	_, err := conn.Write(cmd.Bytes())
	if err != nil {
		return
	}
}

func ParseGetResponse(r io.Reader) (proto.ResponseGet, error) {
	resp := proto.ResponseGet{}
	if err := binary.Read(r, binary.LittleEndian, resp); err != nil {
		log.Println("parse error: ", err)
		return resp, err
	}

	var valueLen int32
	if err := binary.Read(r, binary.LittleEndian, &valueLen); err != nil {
		return resp, err
	}

	resp.Value = make([]byte, valueLen)
	if err := binary.Read(r, binary.LittleEndian, &resp.Value); err != nil {
		return resp, err
	}

	return resp, nil
}

func ParseSetResponse(r io.Reader) (proto.ResponseSet, error) {
	resp := proto.ResponseSet{}
	if err := binary.Read(r, binary.LittleEndian, &resp.Status); err != nil {
		return resp, err
	}

	return resp, nil
}
