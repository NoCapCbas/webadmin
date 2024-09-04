package controllers

import (
	"net/http"
	"testing"
)

func Test_UserProfile_Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := executeRequest(req)
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
