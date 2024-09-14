package model

import (
	mongo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Connection = mongo.Session
type Key = bson.ObjectId

func Open(options ...string) (*mongo.Session, error) {
	conn, err := mongo.Dial(options[1])
	if err != nil {
		return nil, err
	}
	return conn, nil
}
