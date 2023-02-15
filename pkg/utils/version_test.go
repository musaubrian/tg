package utils

import (
	"strings"
	"testing"
)


func TestGetVersion(t *testing.T) {
    got, err := GetVersion()
    if len(got) < 0 {
        t.Error("expected a non empty string got an empty string")
    }

    if strings.Contains(got, "v") != true {
        t.Error("Expected true as the tags are prefixed with 'v', got false")
    }
    if err != nil {
        t.Error("Expected no error returned, got an error")
    }
}
