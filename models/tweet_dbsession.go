package models

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"my-go-microblog-sample/config"
)

var (
	// MaxTweetsToFetch is the maximum tweets that are fetched from the db.
	MaxTweetsToFetch = 50
)

// TweetDBSession handles the connections to the database
type TweetDBSession struct {
	backend *mgo.Session
}

// NewTweetDBSession creates a new database session.
func NewTweetDBSession() *TweetDBSession {
	return &TweetDBSession{backend: config.DBSession().Copy()}
}

// Close closes the database session.
func (dbSession *TweetDBSession) Close() {
	dbSession.backend.Close()
}

// Insert inserts a tweet in the database.
func (dbSession *TweetDBSession) Insert(tweet *Tweet) error {
	return dbSession.collection().Insert(tweet)
}

// Update updates the db with the given tweet
func (dbSession *TweetDBSession) Update(tweet *Tweet) error {
	return dbSession.collection().UpdateId(tweet.ID, tweet)
}

// Query creates a query
func (dbSession *TweetDBSession) Query(query bson.M) *mgo.Query {
	criteria := dbSession.collection().Find(query).Limit(MaxTweetsToFetch)

	return criteria
}

// Delete deletes a single tweet that matches the given query.
func (dbSession *TweetDBSession) Delete(query bson.M) error {
	return dbSession.collection().Remove(query)
}

// DeleteAll deletes all instances of tweet that match the given query.
func (dbSession *TweetDBSession) DeleteAll(query bson.M) (int, error) {
	info, err := dbSession.collection().RemoveAll(query)
	return info.Removed, err
}

func (dbSession *TweetDBSession) database() *mgo.Database {
	return dbSession.backend.DB(config.DefaultDBName)
}

func (dbSession *TweetDBSession) collection() *mgo.Collection {
	return dbSession.database().C("tweets")
}

// vi:syntax=go
