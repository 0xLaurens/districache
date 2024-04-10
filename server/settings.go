package server

type Setting func(*Server)

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
