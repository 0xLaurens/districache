package main

import (
	"github.com/0xlaurens/districache/cache"
	"github.com/0xlaurens/districache/server"
	"log"
)

func main() {
	srv := server.NewServer(cache.New(), server.MakeLeader(true))
	log.Fatal(srv.Run())
}
