package mongo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func Test_DB_Users_GetDetail(t *testing.T) {
	users := Users{}
	users.RefreshSession(db, dbName)

	user, err := users.GetDetail(bson.ObjectIdHex("507f1f77bcf86cd799439011"))
	if err != nil {
		t.Errorf("error getting user detail: %v", err)
	} else if user == nil {
		t.Errorf("user not found")
	} else if user.Email != "test@example.com" {
		t.Errorf("expected user email to be test@example.com, got %s", user.Email)
	}
}

func Test_DB_Users_SignUp(t *testing.T) {
	users := Users{}
	users.RefreshSession(db, dbName)

	acct, err := users.SignUp("test@example.com", "password123")
	if err != nil {
		t.Errorf("error signing up account: %v", err)
	} else if acct.Email != "test@example.com" {
		t.Errorf("expected account email to be test@example.com, got %s", acct.Email)
	} else if len(acct.Users) != 1 {
		t.Errorf("expected account to have 1 user, got %d", len(acct.Users))
	} else if acct.Users[0].Email != "test@example.com" {
		t.Errorf("expected user email to be test@example.com, got %s", acct.Users[0].Email)
	}

}
