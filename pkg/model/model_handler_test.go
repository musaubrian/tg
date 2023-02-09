package model

import "testing"

func TestGetPath(t *testing.T) {
	got := GetPath()

	if got == "" {
		t.Error("Expected a non-empty string got an empty string")
	}
}
