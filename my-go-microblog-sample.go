package main

import (
	"os"

	"github.com/labstack/echo"
	"my-go-microblog-sample/config"
	"my-go-microblog-sample/controllers"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	return port
}

func main() {
	e := echo.New()

	config.Setup(e)
	controllers.Setup(e.Router())

	e.File("/", "index.html")
	e.Static("/static", "static")

	err := e.Start(":" + getPort())
	if err != nil {
		panic(err)
	}
}

// vi:syntax=go
