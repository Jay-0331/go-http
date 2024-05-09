package server

import "strconv"

type Response struct {
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       string
}

const (
	CRLF = "\r\n"
)

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) String() string {
	response := HTTPVersion + " " + strconv.Itoa(r.StatusCode) + " " + r.StatusText + CRLF
	for k, v := range r.Headers {
		response += k + ": " + v + CRLF
	}
	response += CRLF
	response += r.Body
	return response
}
