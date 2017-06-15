package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var storage Storage

func main() {

	app := iris.New()

	commonLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
	})

	app.Use(commonLogger)

	// log http errors should be done manually
	errorLogger := logger.New()

	app.OnErrorCode(iris.StatusNotFound, func(ctx context.Context) {
		errorLogger(ctx)
		msg := errorMessage{iris.StatusNotFound, "Route is not supported"}
		ctx.JSON(msg)
	})

	app.Handle("GET", "/ping", func(ctx context.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.Text("pong")
	})

	app.Handle("GET", "/topics", func(ctx context.Context) {
		getTopicsHandler(ctx)
	})

	app.Handle("POST", "/topics", func(ctx context.Context) {
		addTopicHandler(ctx)
	})

	app.Handle("PUT", "/topics", func(ctx context.Context) {
		updateTopicHandler(ctx)
	})

	config := loadConfig()
	app.Run(iris.Addr(fmt.Sprintf(":%d", config.Listen)), iris.WithConfiguration(iris.Configuration{ // default configuration:
		DisableBanner: true,
	}))
}

func loadConfig() config {
	var filename string

	// register flags
	flag.StringVar(&filename, "config", "", "config filename")
	flag.StringVar(&filename, "c", "", "config filename (shorthand)")

	flag.Parse()

	config := config{}

	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}

	return config
}
