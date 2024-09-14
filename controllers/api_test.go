package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/NoCapCbas/webadmin/data"
)

func logger(next http.Handler) http.Handler {
	// mock logger
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	db := &data.DB{}
	if err := db.Open("unit", "test"); err != nil {
		log.Fatal("unable to connect to the database:", err)
	}
	api := &API{
		DB:     db,
		Logger: logger,
	}

	rec := httptest.NewRecorder()
	api.ServeHTTP(rec, req)
	return rec
}
