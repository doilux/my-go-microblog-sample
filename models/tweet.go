package models

import (
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"my-go-microblog-sample/helpers"
)

var (
	// ErrInvalidDoc is return when the document is not valid
	ErrInvalidDoc = errors.New("document is invalid")
)

// Tweet is a model
type Tweet struct {
	ID        bson.ObjectId  `bson:"_id,omitempty"`
	Body      string         `bson:"body"`
	CreatedAt int64          `bson:"created_at"`
	UpdatedAt int64          `bson:"updated_at"`
	Errors    helpers.Errors `bson:"-"`
}

// NewTweet creates a new instance of Tweet
func NewTweet() *Tweet {
	tweet := &Tweet{}

	// Set default values here

	return tweet
}

// SaveWithSession saves the Tweet using the given db session
func (tweet *Tweet) SaveWithSession(dbSession *TweetDBSession) error {
	isNewRecord := (len(tweet.ID) == 0)
	tweet.setDefaultFields(isNewRecord)

	if !tweet.IsValid() {
		return ErrInvalidDoc
	}

	if isNewRecord {
		return tweet.insert(dbSession)
	}

	return tweet.update(dbSession)
}

// Save saves the Tweet
func (tweet *Tweet) Save() error {
	dbSession := NewTweetDBSession()
	defer dbSession.Close()

	return tweet.SaveWithSession(dbSession)
}

func (tweet *Tweet) DeleteWithSession(dbSession *TweetDBSession) error {
	return dbSession.Delete(bson.M{
		"_id": tweet.ID,
	})
}

func (tweet *Tweet) Delete() error {
	dbSession := NewTweetDBSession()
	defer dbSession.Close()

	return tweet.DeleteWithSession(dbSession)
}

func (tweet *Tweet) setDefaultFields(isNewRecord bool) {
	if isNewRecord {
		tweet.ID = bson.NewObjectId()
		tweet.CreatedAt = time.Now().Unix()
	} else {
		tweet.UpdatedAt = time.Now().Unix()
	}
}

func (tweet *Tweet) insert(dbSession *TweetDBSession) error {
	return dbSession.Insert(tweet)
}

func (tweet *Tweet) update(dbSession *TweetDBSession) error {
	return dbSession.Update(tweet)
}

// ToResponseMap converts the Tweet to a map to be returned by the API as a response.
func (tweet *Tweet) ToResponseMap() helpers.ResponseMap {
	return helpers.ResponseMap{
		"id":   tweet.ID,
		"body": tweet.Body,
	}
}

// ErrorMessages returns validation errors.
func (tweet *Tweet) ErrorMessages() helpers.ResponseMap {
	errorMessages := helpers.ResponseMap{}
	for fieldName, message := range tweet.Errors.Messages {
		errorMessages[fieldName] = message
	}

	return errorMessages
}

// IsValid returns true if all validations are passed.
func (tweet *Tweet) IsValid() bool {
	tweet.Errors.Clear()

	// Run validations here.

	return !tweet.Errors.HasMessages()
}

// SetAttributes sets the attributes for this Tweet
func (tweet *Tweet) SetAttributes(params url.Values) {
	// for avoiding compile unused strconv error.
	var _ = strconv.FormatBool(true)

	if value, ok := params["body"]; ok {
		tweet.Body = value[0]
	}
}

// FindOneTweet finds a tweet in the database using the given `query`
func FindOneTweet(query bson.M) (*Tweet, error) {
	dbSession := NewTweetDBSession()
	defer dbSession.Close()

	criteria := dbSession.Query(query)
	return FindOneTweetByCriteria(criteria)
}

// FindOneTweetByCriteria finds a single document given a criteria
func FindOneTweetByCriteria(criteria *mgo.Query) (*Tweet, error) {
	tweet := &Tweet{}
	err := criteria.One(tweet)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}

// FindOneTweetByID finds a Tweet in the database using the given `id`
func FindOneTweetByID(id string) (*Tweet, error) {
	tweetID := bson.ObjectIdHex(id)
	return FindOneTweet(bson.M{"_id": tweetID})
}

// LoadTweets takes a query a tries to convert it into a Tweet array.
func LoadTweets(query *mgo.Query) ([]*Tweet, error) {
	tweets := []*Tweet{}

	err := query.All(&tweets)
	if err != nil {
		return tweets, err
	}

	return tweets, nil
}

// vi:syntax=go
