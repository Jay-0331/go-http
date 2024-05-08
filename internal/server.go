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

var HTTPVersion = "HTTP/1.1"

func NewServer(port int, host string) *Server {
	server := &Server{
		Port: port, 
		Host: host, 
		Listener: nil,
	}
	return server
}

func (server *Server) Start() {
	// Implement the server start logic here
	listener, err := net.Listen("tcp", server.Host + ":" + strconv.Itoa(server.Port))
	if err != nil {
		fmt.Println("Failed to bind to port ", server.Port)
		os.Exit(1)
	}
	server.Listener = listener
}

func (server *Server) Accept() net.Conn {
	// Implement the server accept logic here
	conn, err := server.Listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	return conn
}

func (server *Server) Close() {
	// Implement the server close logic here
	server.Listener.Close()
}

func (server *Server) Send(conn net.Conn, body string, statusCode int, statusText string, headers map[string]string) {
	// Implement the server send logic here
	req := server.Receive(conn)
	switch req.Path {
	case "/":
		res := NewResponse(statusCode, headers, body, statusText)
		writeResp(conn, res.String())
	default:
		res := NewResponse(404, headers, "Not Found", "Not Found")
		writeResp(conn, res.String())
	}
}

func (server *Server) Receive(conn net.Conn) *Request {
	// Implement the server receive logic here
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		os.Exit(1)
	}
	req := ParseRequest(string(buf[:n]))
	return req
}

func writeResp(conn net.Conn, resp string) {
	_, err := conn.Write([]byte(resp))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		os.Exit(1)
	}
}