package data

import (
	"log"

	"github.com/NoCapCbas/webadmin/data/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SeedDatabase(session *mgo.Session, dbName string) {
	db := session.DB(dbName)
	collection := db.C("users")

	// Create index for email field
	index := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	if err := collection.EnsureIndex(index); err != nil {
		log.Fatalf("Error creating index: %v", err)
	}

	// Seed users
	users := []model.Account{
		{
			ID:    bson.NewObjectId(),
			Email: "alice@example.com",
			Users: []model.User{
				{
					ID:    bson.NewObjectId(),
					Email: "alice@example.com",
				},
			},
		},
		{
			ID:    bson.NewObjectId(),
			Email: "bob@example.com",
			Users: []model.User{
				{
					ID:    bson.NewObjectId(),
					Email: "bob@example.com",
				},
			},
		},
	}

	for _, user := range users {
		if err := collection.Insert(user); err != nil {
			log.Fatalf("Error seeding database: %v", err)
		}
	}
}
