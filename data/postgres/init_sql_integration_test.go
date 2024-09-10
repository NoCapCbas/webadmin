// go:build integration && !mongo

package postgres 

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/NoCapCbas/webadmin/data/model"
)

var db *sql.DB

func TestMain(m *testing.M) {
	conn, err := model.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
	db = conn
	os.Exit(m.Run())
}
