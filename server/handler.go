package server

import (
	"fmt"
	"github.com/0xlaurens/districache/proto"
	"log"
	"net"
	"time"
)

func (s *Server) handleGetCommand(conn net.Conn, cmd *proto.CommandGet) error {
	log.Println("HANDLE GET COMMAND")
	var resp proto.ResponseGet
	val, err := s.cache.Get(cmd.Key)
	if err != nil {
		resp.Status = proto.StatusError
		_, _ = conn.Write(resp.Bytes())
		return nil
	}
	resp.Status = proto.StatusOK
	resp.Value = val

	_, err = conn.Write(resp.Bytes())
	if err != nil {
		log.Println("err", err)
	}

	return nil
}

func (s *Server) handleSetCommand(conn net.Conn, cmd *proto.CommandSet) error {
	fmt.Println("HANDLE SET COMMAND")
	var resp proto.ResponseSet
	if err := s.cache.Set(cmd.Key, cmd.Value, time.Duration(cmd.TTL)); err != nil {
		resp.Status = proto.StatusError
		_, err := conn.Write(resp.Bytes())
		return err
	}

	resp.Status = proto.StatusOK
	_, _ = conn.Write(resp.Bytes())
	return nil
}

func (s *Server) handleDeleteCommand(conn net.Conn, cmd *proto.CommandDelete) error {
	log.Println("HANDLE DELETE COMMAND")
	var resp proto.ResponseDelete
	if err := s.cache.Delete(cmd.Key); err != nil {
		resp.Status = proto.StatusError
	}
	resp.Status = proto.StatusOK

	_, _ = conn.Write(resp.Bytes())
	return nil
}
