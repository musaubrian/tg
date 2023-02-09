package utils

import "testing"

func TestGeneratePassword(t *testing.T) {
	got := GeneratePassword()
	want := 15

	if len(got) != want {
		t.Errorf("Got %q, want: %q", got, want)
	}
}
