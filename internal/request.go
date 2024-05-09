package server

import (
	"bufio"
	"strings"
)

type Request struct {
	Method string
	Path   string
	headers map[string]string
	Body string
	params map[string]string
}

type RequestScanner struct {
	bufio.Scanner
}

func ParseRequest(input string) *Request {
	reader := RequestScanner{
		Scanner: *bufio.NewScanner(strings.NewReader(input)),
	}
	req := &Request{
		Method: "",
		Path: "",
		headers: make(map[string]string),
		Body: "",
		params: make(map[string]string),
	}
	reader.Scan()
	req.ParseRequestLine(reader.Text())
	for reader.Scan() {
		// Implement the request parsing logic here
		if reader.Text() == "" {
			break
		}
		req.ParseRequestHeader(reader.Text())
	}
	return req
}

func (r *Request) ParseRequestLine(line string) {
	parts := strings.Split(line, " ")
	HTTPVersion = parts[2]
	if len(parts) == 3 {
		r.Method = parts[0]
		r.Path = parts[1]
	}
}

func (r *Request) ParseRequestHeader(line string) {
	parts := strings.Split(line, ": ")
	if len(parts) == 2 {
		r.headers[strings.ToLower(parts[0])] = strings.ToLower(parts[1])
	}
}

func (r *Request) AddParam(key, value string) {
	r.params[key] = value
}

func (r *Request) GetParam(key string) string {
	return r.params[key]
}

func (r *Request) GetHeader(key string) string {
	return r.headers[key]
}