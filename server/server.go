package server

import (
	"fmt"
	"github.com/0xlaurens/districache/cache"
	"github.com/0xlaurens/districache/command"
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
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error %s\n", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("conn read error (%s)\n", err)
			break
		}

		go s.handleCommand(conn, buf[:n])
	}
}

func (s *Server) handleCommand(conn net.Conn, rawCmd []byte) {
	cmd, err := command.ParseCmd(rawCmd)
	if err != nil {
		log.Printf("parse cmd error (%s)\n", err)
		_, _ = conn.Write([]byte("invalid command syntax"))
		return
	}
	switch cmd.GetType() {
	case command.CMDGet:
		if err := s.handleGetCmd(conn, cmd.(command.BaseCMD)); err != nil {
			return
		}
		break
	case command.CMDSet:
		if err := s.handleSetCmd(conn, cmd.(command.SetCMD)); err != nil {
			return
		}
		break
	case command.CMDDelete:
		if err := s.handleDeleteCmd(conn, cmd.(command.BaseCMD)); err != nil {
			return
		}
		break
	}

	log.Printf("received cmd (%s)", cmd)
}

func (s *Server) handleSetCmd(conn net.Conn, cmd command.SetCMD) error {
	log.Printf("handling set command: %s\n", cmd)
	err := s.cache.Set(cmd.Key, cmd.Value, cmd.TTL)
	if err != nil {
		return err
	}

	_, _ = conn.Write([]byte("OK"))

	return nil
}

func (s *Server) handleGetCmd(conn net.Conn, cmd command.BaseCMD) error {
	log.Printf("handling GET command: %s\n", cmd)
	res, err := s.cache.Get(cmd.Key)
	if err != nil {
		_, _ = conn.Write([]byte("(nil)"))
		return err
	}
	_, _ = conn.Write(res)

	return nil
}

func (s *Server) handleDeleteCmd(conn net.Conn, cmd command.BaseCMD) error {
	log.Printf("handling DELETE command: %s\n", cmd)
	err := s.cache.Delete(cmd.Key)
	if err != nil {
		return err
	}
	_, _ = conn.Write([]byte("OK"))
	return nil
}
