package main

import (
	"os"
	"strings"

	server "github.com/codecrafters-io/http-server-starter-go/internal"
)

func main() {
	router := server.NewRouter()
	router.GET("/", func(ctx server.Context) string {
		return ctx.Send("", 200, nil)
	})
	router.POST("/files/:filename", func(ctx server.Context) string {
		file, err := os.OpenFile(ctx.GetFilepath() + "/" + ctx.Request.GetParam("filename"), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return ctx.Send("", 500, nil)
		}
		defer file.Close()
		file.Write([]byte(ctx.Request.GetBody()))
		return ctx.Send("", 201, nil)
	})
	router.GET(("/files/:filename"), func(ctx server.Context) string {
		file, err := os.Open(ctx.GetFilepath() + "/" + ctx.Request.GetParam("filename"))
		if err != nil {
			return ctx.Send("", 404, nil)
		}
		defer file.Close()
		fileStat, err := file.Stat()
		if err != nil {
			return ctx.Send("", 500, nil)
		}
		if fileStat.IsDir() {
			return ctx.Send("", 404, nil)
		}
		fileContent := make([]byte, fileStat.Size())
		file.Read(fileContent)
		return ctx.Send(string(fileContent), 200, map[string]string{"Content-Type": "application/octet-stream"})
	})
	router.GET("/echo/:message", func(ctx server.Context) string {
		message := ctx.Request.GetParam("message")
		if strings.Contains(ctx.Request.GetHeader("accept-encoding"), "gzip") {
			ctx.Response.SetHeader("Content-Encoding", "gzip")
		}
		return ctx.Send(message, 200, nil)
	})
	router.GET("/user-agent", func(ctx server.Context) string {
		return ctx.Send(ctx.Request.GetHeader("user-agent"), 200, nil)
	})
	router.GET("*", func(ctx server.Context) string {
		return ctx.Send("", 404, nil)
	})
	server := server.NewServer(4221, "0.0.0.0", router)
	server.Start()
}
