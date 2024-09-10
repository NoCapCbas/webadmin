package data

import (
	"github.com/NoCapCbas/webadmin/data/model"
)

type UserServices interface {
	GetDetail(id model.Key) (*model.User, error) 
}

type DB struct {
	DatabaseName string
	Connection   *model.Connection
	CopySession bool

	Users UserServices
}
