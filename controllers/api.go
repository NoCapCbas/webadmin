package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NoCapCbas/webadmin/engine"
)

// API is starting point of api
// Responsible for routing and handling requests	k
type API struct {
	Logger func(http.Handler) http.Handler
	User   *engine.Route
}

// ServeHTTP is the entry point for all requests
func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, engine.ContextOriginalPath, r.URL.Path)

	var next *engine.Route
	var head string
	head, r.URL.Path = engine.ShiftPath(r.URL.Path)
	if head == "user" {
		next = newUser()
	} else {
		next = newError(fmt.Errorf("path not found: %s", r.URL.Path), http.StatusNotFound)
	}

	if next.Logger {
		next.Handler = api.Logger(next.Handler)
	}

	next.Handler.ServeHTTP(w, r.WithContext(ctx))
}

func newError(err error, statusCode int) *engine.Route {
	return &engine.Route{
		Logger: true,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			engine.Respond(w, r, statusCode, err)
		}),
	}
}

func newUser() *engine.Route {
	return &engine.Route{
		Logger: true,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			engine.Respond(w, r, http.StatusOK, "User")
		}),
	}
}

func NewAPI() *API {
	return &API{
		Logger: engine.Logger,
	}
}
