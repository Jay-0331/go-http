package server

import (
	"bufio"
	"strings"
)

type Request struct {
	Method string
	Path   string
	Headers map[string]string
	Body string
}

type RequestScanner struct {
	bufio.Scanner
}

func ParseRequest(req string) *Request {
	reader := RequestScanner{
		Scanner: *bufio.NewScanner(strings.NewReader(req)),
	}
	for reader.Scan() {
		// Implement the request parsing logic here
		line := reader.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, " ")
		if len(parts) == 3 {
			HTTPVersion = parts[2]
			return &Request{
				Method: parts[0],
				Path: parts[1],
			}
		}
	}
	return nil
}