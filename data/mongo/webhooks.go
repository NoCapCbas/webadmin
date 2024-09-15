//go:build !mem

package mongo

import (
	"context"
	"strings"
	"time"

	"github.com/NoCapCbas/webadmin/data/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Webhooks struct {
	DB *mongo.Database
}

// Add inserts a new webhook
func (wh *Webhooks) Add(accountID model.Key, events, url string) error {
	var hooks []interface{}
	en := strings.Split(events, ",")
	for _, e := range en {
		hooks = append(hooks, model.Webhook{
			ID:        primitive.NewObjectID(),
			AccountID: accountID,
			EventName: strings.Trim(e, " "),
			TargetURL: url,
			IsActive:  true,
			Created:   time.Now(),
		})
	}
	_, err := wh.DB.Collection("webhooks").InsertMany(context.TODO(), hooks)
	return err
}

// Delete removes all matching target webhook url
func (wh *Webhooks) Delete(accountID model.Key, event, url string) error {
	filter := bson.M{"accountId": accountID, "event": event, "url": url}
	_, err := wh.DB.Collection("webhooks").DeleteMany(context.TODO(), filter)
	return err
}

// List returns the webhook entries for an account
func (wh *Webhooks) List(accountID model.Key) ([]model.Webhook, error) {
	var results []model.Webhook
	filter := bson.M{"accountId": accountID}
	cursor, err := wh.DB.Collection("webhooks").Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// AllSubscriptions returns all webhooks for an event
func (wh *Webhooks) AllSubscriptions(event string) ([]model.Webhook, error) {
	var results []model.Webhook
	filter := bson.M{"event": event}
	cursor, err := wh.DB.Collection("webhooks").Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (wh *Webhooks) RefreshSession(client *mongo.Client, dbName string) {
	wh.DB = client.Database(dbName)
}

// Close is no longer needed with the new driver
func (wh *Webhooks) Close() {
	// No-op
}
