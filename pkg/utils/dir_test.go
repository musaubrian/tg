package utils

import "testing"

func TestCreateDir(t *testing.T) {
	got := CreateDir()
	if got != nil {
		t.Error("Expected nil got an error")
	}
}
