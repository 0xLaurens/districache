package server

import (
	"fmt"
	"github.com/0xlaurens/districache/cache"
	"log"
	"net"
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
		_, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error %s\n", err)
			continue
		}
	}

	return nil
}
