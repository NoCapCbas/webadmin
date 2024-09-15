//go:build !mem

package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection = *mongo.Client
type Key = primitive.ObjectID

func Open(args ...string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(args[1]))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func keyToString(id Key) string {
	return id.Hex()
}
