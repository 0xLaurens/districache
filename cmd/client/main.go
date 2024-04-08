package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return
	}
	sendCommand(conn, "GET hello")
	readResponse(conn)

	sendCommand(conn, "SET hello world 25000")
	readResponse(conn)

	sendCommand(conn, "GET hello")
	readResponse(conn)

	sendCommand(conn, "SET hello bozo")
	readResponse(conn)

	sendCommand(conn, "GET hello")
	readResponse(conn)

	sendCommand(conn, "DEL hello")
	readResponse(conn)

	sendCommand(conn, "GET hello")
	readResponse(conn)
}

func readResponse(conn net.Conn) {
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("read error (%s)\n", err)
		return
	}
	res := buf[:n]
	log.Println(string(res))
}

func sendCommand(conn net.Conn, cmd string) {
	_, err := conn.Write([]byte(cmd))
	if err != nil {
		log.Printf("write error (%s)\n", err)
		return
	}
}
