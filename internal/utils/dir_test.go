package utils_test

import (
	"testing"

	"github.com/musaubrian/tinygo/internal/utils"
)

func TestCreateDir(t *testing.T) {
	if utils.CreateDir() != nil {
		t.Error("Expected nil got an error")
	}
}

func TestGetPath(t *testing.T) {
	fullPath, err := utils.GetPath()
	if err != nil {
		t.Error("Expected nil got an error")
	}
	if len(fullPath) < 0 {
		t.Error("Got an empty string")
	}
}
