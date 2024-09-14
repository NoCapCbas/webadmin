package data

import (
	"github.com/NoCapCbas/webadmin/data/model"
)

type SessionRefresher interface {
	RefreshSession(db *model.Connection, dbName string)
}

type UserServices interface {
	SessionRefresher
	GetDetail(id model.Key) (*model.User, error)
}

type DB struct {
	DatabaseName string
	Connection   *model.Connection
	CopySession  bool

	Users UserServices
}
