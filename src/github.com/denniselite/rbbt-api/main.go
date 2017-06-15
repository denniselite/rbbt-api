package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()
	app.Handle("GET", "/ping", func(ctx context.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.Text("pong")
	})
	app.Run(iris.Addr(":3000"))
}