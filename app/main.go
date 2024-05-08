package main

import (
	"fmt"

	server "github.com/codecrafters-io/http-server-starter-go/internal"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	server := server.NewServer(4221, "0.0.0.0")
	server.Start()
	_ = server.Accept()
}
