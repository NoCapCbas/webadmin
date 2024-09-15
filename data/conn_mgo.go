//go:build !mem

package data

import (
	"fmt"
	"os"

	"github.com/NoCapCbas/webadmin/data/model"
	"github.com/NoCapCbas/webadmin/data/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) Open(driverName, dataSourceName string) error {
	fmt.Println("inside mongo sourceName", dataSourceName)
	client, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	fmt.Println("mongo connection successful")

	db.DatabaseName = os.Getenv("DB_NAME")
	database := client.Database(db.DatabaseName)

	db.Users = &mongo.Users{DB: database}
	db.Webhooks = &mongo.Webhooks{DB: database}

	db.Connection = client

	// CopySession is no longer needed with the new driver
	// db.CopySession = true
	return nil
}
