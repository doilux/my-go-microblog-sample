package controllers

import (
	"github.com/labstack/echo"
	"my-go-microblog-sample/helpers"
	"my-go-microblog-sample/models"
)

func findTweetByID(c echo.Context) (*models.Tweet, *helpers.ResponseError) {
	c.Request().ParseForm()

	tweet, _ := models.FindOneTweetByID(c.Param("tweet_id"))
	if tweet == nil {
		return nil, ErrTweetNotFound
	}

	return tweet, nil
}

// vi:syntax=go
