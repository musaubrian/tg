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
