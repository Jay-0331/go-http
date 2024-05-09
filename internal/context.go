package server

type Context struct {
	Request  *Request
	Response *Response
	filepath string
}

func NewContext() *Context {
	return &Context{
		Request: nil,
		Response: NewResponse(),
	}
}

func (c *Context) Send(body string, statusCode int, statusText string, headers map[string]string) string {
	c.Response.StatusCode = statusCode
	c.Response.StatusText = statusText
	c.Response.Headers = headers
	c.Response.Body = body
	return c.Response.String()
}

func (c *Context) SetRequest(request *Request) {
	c.Request = request
}

func (c *Context) SetSatusCode(statusCode int) {
	c.Response.StatusCode = statusCode
}

func (c *Context) SetFilepath(filepath string) {
	c.filepath = filepath
}

func (c *Context) GetFilepath() string {
	return c.filepath
}