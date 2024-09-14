//go:build mem

package mem

import (
	"errors"

	"github.com/NoCapCbas/webadmin/data/model"
)

type Users struct {
	store []model.User
}

func (u *Users) GetDetail(id model.Key) (*model.User, error) {
	var user model.User
	for _, usr := range u.store {
		if usr.ID == id {
			user = usr
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (u *Users) RefreshSession(conn *bool, dbName string) {
	u.store = append(u.store, model.User{
		ID: 1,
		Email: "test@test.com",
	}
}
