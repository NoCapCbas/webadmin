package data

import (
	"github.com/NoCapCbas/webadmin/data/model"
)

func (db *DB) Open(driverName, dataSource string) error {
	conn, err := model.Open(driverName, dataSource)
	if err != nil {
		return err
	}

	db.CopySession = true
	db.Connection = conn
	return nil
}
