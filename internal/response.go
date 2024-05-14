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
	return &Response{
		Headers: make(map[string]string),
	}
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

func (r *Response) SetHeader(key, value string) {
	r.Headers[key] = value
}

func (r *Response) SetHeaders(headers map[string]string) {
	for key, value := range headers {
		r.Headers[key] = value
	}
}

func (r *Response) SetBody(body string) {
	if r.Headers["Content-Type"] == "" {
		r.SetHeader("Content-Type", "text/plain")
	}
	r.SetHeader("Content-Length", strconv.Itoa(len(body)))
	r.Body = body
}