//go:build !mongo

package data

import (
  "log"
	"github.com/NoCapCbas/webadmin/data/model"
  "github.com/NoCapCbas/webadmin/data/postgres"
)

func (db *DB) Open(driverName, dataSourceName string) error {
	conn, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
  // initialize services 
  db.Users = &postgres.Users{DB: conn}
  if db.Users == nil {
    log.Fatal("Failed to initialize UserServices")
  }

	// initialize the database
	db.Connection = conn
	return nil
}
