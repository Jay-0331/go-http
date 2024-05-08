package main

import (
	server "github.com/codecrafters-io/http-server-starter-go/internal"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	server := server.NewServer(4221, "0.0.0.0")
	server.Start()
	conn := server.Accept()
	server.Send(conn, "", 200, "OK", nil)
}
