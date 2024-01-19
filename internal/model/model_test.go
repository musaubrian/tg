package model_test

import (
	"testing"

	"github.com/musaubrian/tg/internal/model"
)

func TestSetupDb(t *testing.T) {

	err := model.SetupDB("../../test.db")
	if err != nil {
		t.Error("Expected nil got an error")
	}
}
