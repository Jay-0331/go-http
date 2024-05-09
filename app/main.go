package main

import (
	"fmt"
	"strconv"

	server "github.com/codecrafters-io/http-server-starter-go/internal"
)

func main() {
	router := server.NewRouter()
	router.GET("/", func(ctx server.Context) string {
		fmt.Println(ctx.Request.Path)
		return ctx.Send("", 200, "OK", nil)
	})
	router.GET("/echo/:message", func(ctx server.Context) string {
		fmt.Println(ctx.Request.GetParam("message"))
		message := ctx.Request.GetParam("message")
		return ctx.Send(message, 200, "OK", map[string]string{"Content-Type": "text/plain", "content-length": strconv.Itoa(len(message))})
	})
	router.GET("*", func(ctx server.Context) string {
		fmt.Println(ctx.Request.Path, "not found")
		return ctx.Send("", 404, "Not Found", nil)
	})
	server := server.NewServer(4221, "0.0.0.0", router)
	server.Start()
}
