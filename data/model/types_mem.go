//go:build mem

package model

import "fmt"

type Connection = bool
type Key = int

func Open(args ...string) (bool, error) {
	return true, nil
}

func keyToString(id Key) string {
	return fmt.Sprintf("%d", id)
}
