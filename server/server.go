package server

import (
	"fmt"
	"github.com/0xlaurens/districache/cache"
	"github.com/0xlaurens/districache/proto"
	"io"
	"log"
	"net"
	"time"
)

type Setting func(*Server)

type Server struct {
	port     int
	host     string
	isLeader bool
	cache    cache.Cacher
}

func WithPort(port int) Setting {
	return func(s *Server) {
		s.port = port
	}
}

func WithHost(host string) Setting {
	return func(s *Server) {
		s.host = host
	}
}

func MakeLeader(leader bool) Setting {
	return func(s *Server) {
		s.isLeader = leader
	}
}

func NewServer(cache cache.Cacher, settings ...Setting) *Server {
	server := &Server{
		cache: cache,
		host:  "127.0.0.1",
		port:  3000,
	}

	for _, setting := range settings {
		setting(server)
	}

	return server
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		return fmt.Errorf("listening error %s", err)
	}

	log.Printf("Server started on (%s:%d)\n", s.host, s.port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error %s\n", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	for {
		cmd, err := proto.ParseCommand(conn)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("parse command error:", err)
			break
		}
		go s.handleCommand(conn, cmd)
	}
}

func (s *Server) handleCommand(conn net.Conn, cmd any) {
	switch v := cmd.(type) {
	case *proto.CommandGet:
		_ = s.handleGetCommand(conn, v)
	case *proto.CommandSet:
		_ = s.handleSetCommand(conn, v)
	case *proto.CommandDelete:
		_ = s.handleDeleteCommand(conn, v)
	}
}

func (s *Server) handleGetCommand(conn net.Conn, cmd *proto.CommandGet) error {
	log.Println("HANDLE GET COMMAND")
	var resp proto.ResponseGet
	val, err := s.cache.Get(cmd.Key)
	if err != nil {
		resp.Status = proto.StatusNone
		resp.Value = []byte("(nil)")
		_, err := conn.Write(resp.Bytes())
		return err
	}
	resp.Status = proto.StatusOK
	resp.Value = val

	_, _ = conn.Write(resp.Bytes())

	return nil
}

func (s *Server) handleSetCommand(conn net.Conn, cmd *proto.CommandSet) error {
	log.Println("HANDLE SET COMMAND")
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
