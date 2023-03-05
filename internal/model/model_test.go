package model_test

import (
	"testing"

	"github.com/musaubrian/tinygo/internal/model"
)

func TestSetupDb(t *testing.T) {
	err := model.SetupDB()
	if err != nil {
		t.Error("Expected nil got an error")
	}
}
