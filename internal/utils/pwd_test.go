package utils_test

import (
	"testing"

	"github.com/atotto/clipboard"
	"github.com/musaubrian/tinygo/internal/utils"
)

func TestGeneratePassword(t *testing.T) {
	got := utils.GeneratePassword()
	want := 15

	if len(got) != want {
		t.Errorf("Got %q, want: %q", got, want)
	}
}

func TestCopyToClipboard(t *testing.T) {
	testStr := "This should be copied and retrieved from the clipboard"
	err := utils.CopyToClipboard(testStr)
	if err != nil {
		t.Errorf("Expected nil got %v", err)
	}

	res, err := clipboard.ReadAll()
	if err != nil {
		t.Errorf("Expected nil got %v", err)
	}
	if res != testStr {
		t.Errorf("Value copied is not same as what was copied, got %s expected %s", testStr, res)
	}
}
