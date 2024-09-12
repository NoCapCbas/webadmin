package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NoCapCbas/webadmin/data"
	"github.com/NoCapCbas/webadmin/engine"
)

type User struct{}

func newUser() *engine.Route {
	var u interface{} = User{}
	return &engine.Route{
		Logger:  true,
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
	engine.Respond(w, r, http.StatusOK, "viewing profile")
}

func (u User) detail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// id := ctx.Value(engine.ContextUserID).(int64)
	// db := ctx.Value(engine.ContextDatabase).(*data.DB)
	var ok bool

	// Retrieve the user ID from the context
	idCtx := ctx.Value(engine.ContextUserID)
	if idCtx == nil {
		log.Println("User ID not found in context")
		engine.Respond(w, r, http.StatusBadRequest, "Missing user ID")
		return
	}

	// Assert that the userID is an int64
	id, ok := idCtx.(int64)
	if !ok {
		log.Printf("Invalid type for user ID: expected int64, got %T\n", id)
		engine.Respond(w, r, http.StatusBadRequest, "Invalid user ID type")
		return
	}

	// Retrieve the database from the context
	dbCtx := ctx.Value(engine.ContextDatabase)
	if dbCtx == nil {
		log.Println("Database not found in context")
		engine.Respond(w, r, http.StatusInternalServerError, "Database connection error")
		return
	}

	// Assert that the db is of type *data.DB
	db, ok := dbCtx.(*data.DB)
	if !ok {
		engine.Respond(w, r, http.StatusInternalServerError, "Internal Server Error")
		log.Fatalf("Invalid type for database connection: expected *data.DB, got %T\n", db)
	}

	var result = new(struct {
		ID    int64     `json:"user_id"`
		Email string    `json:"email"`
		Time  time.Time `json:"time"`
	})

	user, err := db.Users.GetDetail(id)
	if err != nil {
		log.Println("User details not found in db")
		engine.Respond(w, r, http.StatusInternalServerError, err)
		return
	}

	result.ID = user.ID
	result.Email = user.Email
	result.Time = time.Now()

	engine.Respond(w, r, http.StatusOK, result)
}
