package utils_test

import (
	"testing"

	"github.com/musaubrian/tinygo/internal/utils"
)

func TestGeneratePassword(t *testing.T) {
	got := utils.GeneratePassword()
	want := 15

	if len(got) != want {
		t.Errorf("Got %q, want: %q", got, want)
	}
}
