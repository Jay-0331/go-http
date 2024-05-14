package server

type Context struct {
	Request  *Request
	Response *Response
	filepath string
}

var code_map = map[int]string{
	200: "OK",
	201: "Created",
	404: "Not Found",
	500: "Internal Server Error",
}

func NewContext() *Context {
	return &Context{
		Request: nil,
		Response: NewResponse(),
	}
}

func (c *Context) Send(body string, statusCode int, headers map[string]string) string {
	c.Response.StatusCode = statusCode
	c.Response.StatusText = code_map[statusCode]
	c.Response.SetBody(body)
	c.Response.SetHeaders(headers)
	return c.Response.String()
}

func (c *Context) SetRequest(request *Request) {
	c.Request = request
}

func (c *Context) SetSatusCode(statusCode int) {
	c.Response.StatusCode = statusCode
	c.Response.StatusText = code_map[statusCode]
}

func (c *Context) SetFilepath(filepath string) {
	c.filepath = filepath
}

func (c *Context) GetFilepath() string {
	return c.filepath
}