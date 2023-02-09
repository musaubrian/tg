package utils

import "testing"

func TestCreateDir(t *testing.T) {
	if CreateDir() != nil {
		t.Error("Expected nil got an error")
	}
}
