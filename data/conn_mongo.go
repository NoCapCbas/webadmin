//go:build mongo

package data

import (
	"github.com/NoCapCbas/webadmin/data/model"
)

func (db *DB) Open(driverName, dataSourceName string) error {
	conn, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	// copy the session
	db.CopySession = true
	// initialize the database
	db.Connection = conn
	return nil
}
