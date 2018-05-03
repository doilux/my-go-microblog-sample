package controllers

import (
	"github.com/labstack/echo"
)

// Setup sets up all controllers.
func Setup(router *echo.Router) {
	tweets := TweetsController{Router: router}
	tweets.Setup()
}

// vi:syntax=go
