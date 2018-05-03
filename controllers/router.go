package controllers

import (
	"github.com/labstack/echo"
)

// Setup sets up all controllers.
func Setup(router *echo.Router) {
	// ここがscaffoldで追加された部分。
	tweets := TweetsController{Router: router}
	tweets.Setup()
}

// vi:syntax=go
