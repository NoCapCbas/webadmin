package controllers

import (
	"net/http"
	"time"
	"fmt"

	"github.com/NoCapCbas/webadmin/engine"
	"github.com/NoCapCbas/webadmin/data"
)

type User struct{}

func newUser() *engine.Route {
  var u interface{} = User{} 
	return &engine.Route{
		Logger: true,
		Handler: u.(http.Handler),	
  }
}

func (u User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "profile" {
		u.profile(w, r)
		return
	} else if head == "detail" {
		u.detail(w, r)
		return
	}
	newError(fmt.Errorf("path not found"), http.StatusNotFound).Handler.ServeHTTP(w, r)
}

func (u User) profile(w http.ResponseWriter, r *http.Request) {
	engine.Respond(w, r, http.StatusOK, "viewing detail")
}

func (u User) detail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ctx.Value(engine.ContextUserID).(int64)
	db := ctx.Value(engine.ContextDatabase).(*data.DB)

	var result = new(struct {
		ID int64 `json:"user_id"`
		Email string `json:"email"`
		Time time.Time `json:"time"`	
	})

	user, err := db.Users.GetDetail(id)
	if err != nil {
		engine.Respond(w, r, http.StatusInternalServerError, err)
		return
	}

	result.ID = user.ID	
	result.Email = user.Email
	result.Time = time.Now()

	engine.Respond(w, r, http.StatusOK, result)
}
