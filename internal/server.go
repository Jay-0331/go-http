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
	Router
}

var HTTPVersion = "HTTP/1.1"

func NewServer(port int, host string, router Router) *Server {
	server := &Server{
		Port: port, 
		Host: host, 
		Listener: nil,
		Router: router,
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
	for {
		conn := server.Accept()
		go func(conn net.Conn) {
			defer conn.Close()
			ctx := NewContext()
			ctx.Request = server.Receive(conn)
			for _, handlers := range server.Router.MatchRoute(ctx.Request.Method, ctx.Request.Path) {
				for key, value := range handlers.params {
					ctx.Request.AddParam(key, value)
				}
				resp := handlers.handler(*ctx)
				if resp != "" {
					conn.Write([]byte(resp))
					break
				}
			}
		}(conn)
	}
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