package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type Server struct {
	Port int
	Host string
	net.Listener
}

func NewServer(port int, host string) *Server {
	server := &Server{Port: port, Host: host, Listener: nil}
	return server
}

func (s *Server) Start() {
	// Implement the server start logic here
	listener, err := net.Listen("tcp", s.Host + ":" + strconv.Itoa(s.Port))
	if err != nil {
		fmt.Println("Failed to bind to port ", s.Port)
		os.Exit(1)
	}
	s.Listener = listener
}

func (s *Server) Accept() net.Conn {
	// Implement the server accept logic here
	conn, err := s.Listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	return conn
}