package main

import (
	"github.com/0xlaurens/districache/cache"
	"github.com/0xlaurens/districache/server"
	"log"
	"net"
)

func main() {
	srv := server.NewServer(cache.New(), server.MakeLeader(true))

	go func() {
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			return
		}

		_, err = conn.Write([]byte("DELETE hello"))
		if err != nil {
			return
		}
	}()

	log.Fatal(srv.Run())
}
