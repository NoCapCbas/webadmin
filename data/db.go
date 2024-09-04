package data

import (
	"github.com/NoCapCbas/webadmin/data/model"
)

type DB struct {
	DatabaseName string
	Connection   *model.Connection
}
