package postgres

import (
	"database/sql"

	"github.com/NoCapCbas/webadmin/data/model"
)

type Users struct {
	DB *sql.DB
}

func (u *Users) GetDetail(id model.Key) (*model.User, error) {

	var user model.User
	if err := u.DB.QueryRow("SELECT * FROM users WHERE user_id = $1", id).Scan(&user.ID, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil	
}
