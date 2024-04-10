package server

import (
	"fmt"
	"github.com/0xlaurens/districache/cache"
	"github.com/0xlaurens/districache/proto"
	"io"
	"log"
	"net"
)

type Server struct {
	port     int
	host     string
	isLeader bool
	cache    cache.Cacher
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
	case *proto.CommandSet:
		_ = s.handleSetCommand(conn, v)
	case *proto.CommandGet:
		_ = s.handleGetCommand(conn, v)
	case *proto.CommandDelete:
		_ = s.handleDeleteCommand(conn, v)
	}
}
