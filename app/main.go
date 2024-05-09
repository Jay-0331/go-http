package main

import (
	"os"
	"strconv"

	server "github.com/codecrafters-io/http-server-starter-go/internal"
)

func main() {
	router := server.NewRouter()
	router.GET("/", func(ctx server.Context) string {
		return ctx.Send("", 200, "OK", nil)
	})
	router.POST("/files/:filename", func(ctx server.Context) string {
		file, err := os.OpenFile(ctx.GetFilepath() + "/" + ctx.Request.GetParam("filename"), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return ctx.Send("", 500, "Internal Server Error", nil)
		}
		defer file.Close()
		file.Write([]byte(ctx.Request.GetBody()))
		return ctx.Send("", 201, "Created", nil)
	})
	router.GET(("/files/:filename"), func(ctx server.Context) string {
		file, err := os.Open(ctx.GetFilepath() + "/" + ctx.Request.GetParam("filename"))
		if err != nil {
			return ctx.Send("", 404, "Not Found", nil)
		}
		defer file.Close()
		fileStat, err := file.Stat()
		if err != nil {
			return ctx.Send("", 500, "Internal Server Error", nil)
		}
		if fileStat.IsDir() {
			return ctx.Send("", 404, "Not Found", nil)
		}
		fileContent := make([]byte, fileStat.Size())
		file.Read(fileContent)
		return ctx.Send(string(fileContent), 200, "OK", map[string]string{"Content-Type": "application/octet-stream", "content-length": strconv.Itoa(len(fileContent))})
	})
	router.GET("/echo/:message", func(ctx server.Context) string {
		message := ctx.Request.GetParam("message")
		return ctx.Send(message, 200, "OK", map[string]string{"Content-Type": "text/plain", "content-length": strconv.Itoa(len(message))})
	})
	router.GET("/user-agent", func(ctx server.Context) string {
		return ctx.Send(ctx.Request.GetHeader("user-agent"), 200, "OK", map[string]string{"Content-Type": "text/plain", "content-length": strconv.Itoa(len(ctx.Request.GetHeader("user-agent")))})
	})
	router.GET("*", func(ctx server.Context) string {
		return ctx.Send("", 404, "Not Found", nil)
	})
	server := server.NewServer(4221, "0.0.0.0", router)
	server.Start()
}
