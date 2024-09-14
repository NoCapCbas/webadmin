package mongo

import (
	"log"
	"os"
	"testing"

	"github.com/NoCapCbas/webadmin/data/model"
	mongo "gopkg.in/mgo.v2"
)

const (
	dbName = "savvi"
)

var (
	db *mongo.Session
)

func TestMain(m *testing.M) {
	conn, err := model.Open("mongo", "mongodb://localhost:27017")
	if err != nil {
		log.Println("Error opening connection to mongo", err)
		os.Exit(1)
	}
	db = conn
	defer conn.Close()
	os.Exit(m.Run())
}
