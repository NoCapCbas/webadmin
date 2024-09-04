package engine

import (
	"testing"
)

func Test_URL_ShiftPath(t *testing.T) {

	paths := []string{
		"", "", "/",
		"/", "", "/",
		"/test", "test", "/",
		"/test/", "test", "/",
		"/test/1", "test", "/1",
	}

	for i := 0; i < len(paths); i += 3 {
		head, tail := ShiftPath(paths[i])
		if head != paths[i+1] {
			t.Errorf("ShiftPath(%q) = %q, %q, want %q, %q", paths[i], "expected head", paths[i+1], "received", head)
		}
		if tail != paths[i+2] {
			t.Errorf("ShiftPath(%q) = %q, %q, want %q, %q", paths[i], "expected tail", paths[i+2], "received", tail)
		}
	}

}
