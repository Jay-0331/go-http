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
	params map[string]string
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
				Headers: make(map[string]string),
				Body: "",
				params: make(map[string]string),
			}
		}
	}
	return nil
}

func (r *Request) AddParam(key, value string) {
	r.params[key] = value
}

func (r *Request) GetParam(key string) string {
	return r.params[key]
}