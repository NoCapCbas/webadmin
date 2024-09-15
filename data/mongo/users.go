//go:build !mem

package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/NoCapCbas/webadmin/data/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	DB *mongo.Database
}

func (u *Users) SignUp(email, password string) (*model.Account, error) {
	accountID := primitive.NewObjectID()

	acct := model.Account{ID: accountID, Email: email}
	acct.Users = append(acct.Users, model.User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Password: password,
		Token:    model.NewToken(accountID),
		Role:     model.RoleAdmin,
	})
	_, err := u.DB.Collection("users").InsertOne(context.TODO(), acct)
	if err != nil {
		return nil, err
	}
	return &acct, nil
}

func (u *Users) AddToken(accountID, userID model.Key, name string) (*model.AccessToken, error) {
	tok := model.AccessToken{
		ID:    primitive.NewObjectID(),
		Name:  name,
		Token: model.NewToken(accountID),
	}

	filter := bson.M{"_id": accountID, "users._id": userID}
	update := bson.M{"$push": bson.M{"users.$.pat": tok}}
	_, err := u.DB.Collection("users").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return &tok, nil
}

func (u *Users) RemoveToken(accountID, userID, tokenID model.Key) error {
	where := bson.M{"_id": accountID, "users._id": userID}
	update := bson.M{"$pull": bson.M{"users.$.pat": bson.M{"_id": tokenID}}}
	_, err := u.DB.Collection("users").UpdateOne(context.TODO(), where, update)
	return err
}

func (u *Users) Auth(accountID, token string, pat bool) (*model.Account, *model.User, error) {

	id, err := primitive.ObjectIDFromHex(accountID)
	if err != nil {
		return nil, nil, fmt.Errorf("this account id is invalid %s", accountID)
	}

	fmt.Println("looking account", id)

	acct, err := u.GetDetail(id)
	if err != nil {
		fmt.Println("cannot find acct", err)
		return nil, nil, err
	}

	var user model.User
	for _, usr := range acct.Users {
		if pat {
			for _, at := range usr.AccessTokens {
				if at.Token == token {
					user = usr
					break
				}
			}
		} else {
			fmt.Println("comparing", token, "with", usr.Token)
			if usr.Token == token {
				user = usr
				break
			}
		}
	}

	if len(user.Email) == 0 {
		return nil, nil, fmt.Errorf("unable to find this token %s", token)
	}

	return acct, &user, nil
}

func (u *Users) GetDetail(id model.Key) (*model.Account, error) {
	var acct model.Account
	filter := bson.M{"_id": id}
	err := u.DB.Collection("users").FindOne(context.TODO(), filter).Decode(&acct)
	if err != nil {
		return nil, err
	}
	return &acct, nil
}

func (u *Users) GetByStripe(stripeID string) (*model.Account, error) {
	var acct model.Account
	where := bson.M{"stripeId": stripeID}
	if err := u.DB.Collection("users").FindOne(context.TODO(), where).Decode(&acct); err != nil {
		return nil, err
	}
	return &acct, nil
}

// SetSeats set the paid seat for an account
func (u *Users) SetSeats(id model.Key, seats int) error {
	set := bson.M{"$set": bson.M{"seats": seats}}
	where := bson.M{"_id": id}
	_, err := u.DB.Collection("users").UpdateOne(context.TODO(), where, set)
	return err
}

// ConvertToPaid set an account as a paying customer
func (u *Users) ConvertToPaid(id model.Key, stripeID, subID, plan string, yearly bool, seats int) error {
	update := bson.M{"$set": bson.M{
		"stripeId":    stripeID,
		"subId":       subID,
		"plan":        plan,
		"isYearly":    yearly,
		"seats":       seats,
		"subscribed":  time.Now(),
		"trial.trial": false,
	}}
	_, err := u.DB.Collection("users").UpdateOne(context.TODO(), id, update)
	return err
}

// ChangePlan updates an account plan info
func (u *Users) ChangePlan(id model.Key, plan string, yearly bool) error {
	set := bson.M{"$set": bson.M{
		"plan":     plan,
		"isYearly": yearly,
	}}
	_, err := u.DB.Collection("users").UpdateOne(context.TODO(), id, set)
	return err
}

func (u *Users) Cancel(id model.Key) error {
	set := bson.M{"$set": bson.M{
		"subId":    "",
		"plan":     "",
		"isYearly": false,
		"seats":    0,
	}}
	_, err := u.DB.Collection("users").UpdateOne(context.TODO(), id, set)
	return err
}

func (u *Users) RefreshSession(client *mongo.Client, dbName string) {
	u.DB = client.Database(dbName)
}

func (u *Users) Close() {
	// No need to close explicitly in the new driver
}
