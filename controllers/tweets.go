package controllers

import (
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"my-go-microblog-sample/helpers"
	"my-go-microblog-sample/models"
	"log"
)

var (
	// ErrTweetNotFound is returned when the tweet was not found.
	ErrTweetNotFound = helpers.NewResponseError(404, "tweet not found")

	// ErrCannotDeleteTweet is returned when the tweet can't be deleted.
	ErrCannotDeleteTweet = helpers.NewResponseError(404, "cannot delete tweet")
)

// TweetsController is the Tweet controller.
type TweetsController struct {
	Router *echo.Router
}

func (controller *TweetsController) createTweet(c echo.Context) error {
	c.Request().ParseForm()

	// logging form
	log.Print(c.Request().Form)

	tweet := &models.Tweet{}
	tweet.SetAttributes(c.Request().Form)

	tweet.Save()
	if len(tweet.ErrorMessages()) == 0 {
		return helpers.JSONResponseObject(c, 200, tweet)
	}
	return helpers.JSONResponse(c, 400, tweet.ErrorMessages())
}

func (controller *TweetsController) listTweets(c echo.Context) error {
	session := models.NewTweetDBSession()
	defer session.Close()

	query := session.Query(bson.M{})
	query.Sort("-created_at").Limit(20)

	tweets, _ := models.LoadTweets(query)
	tweetsResponse := make([]helpers.ResponseMap, len(tweets))
	for i, tweets := range tweets {
		tweetsResponse[i] = tweets.ToResponseMap()
	}

	return helpers.JSONResponseArray(c, 200, tweetsResponse)
}

func (controller *TweetsController) getTweet(c echo.Context) error {
	tweet, err := findTweetByID(c)

	if err != nil {
		return helpers.JSONResponseError(c, err)
	}

	return helpers.JSONResponseObject(c, 200, tweet)
}

func (controller *TweetsController) updateTweet(c echo.Context) error {
	tweet, err := findTweetByID(c)

	if err != nil {
		return helpers.JSONResponseError(c, err)
	}

	tweet.SetAttributes(c.Request().Form)

	tweet.Save()
	if len(tweet.ErrorMessages()) == 0 {
		return helpers.JSONResponseObject(c, 200, tweet)
	}
	return helpers.JSONResponse(c, 400, tweet.ErrorMessages())
}

func (controller *TweetsController) deleteTweet(c echo.Context) error {
	tweet, err := findTweetByID(c)

	if err != nil {
		return helpers.JSONResponseError(c, err)
	}

	if err := tweet.Delete(); err == nil {
		return helpers.JSONResponse(c, 200, helpers.ResponseMap{})
	}

	return helpers.JSONResponseError(c, ErrCannotDeleteTweet)
}

// Setup sets up routes for the Tweet controller.
func (controller *TweetsController) Setup() {
	controller.Router.Add("POST", "/tweets", controller.createTweet)
	controller.Router.Add("GET", "/tweets", controller.listTweets)
	controller.Router.Add("GET", "/tweets/:tweet_id", controller.getTweet)
	controller.Router.Add("PUT", "/tweets/:tweet_id", controller.updateTweet)
	controller.Router.Add("DELETE", "/tweets/:tweet_id", controller.deleteTweet)
}

// vi:syntax=go
