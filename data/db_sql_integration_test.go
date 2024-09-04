//go:build integration && !mongo

package data

import (
	"testing"
)

func Test_DB_Open(t *testing.T) {
	db := DB{}
	if err := db.Open("postgres", ""); err != nil {
		t.Fatalf("unable to connect to database: %v", err)
	}
}
