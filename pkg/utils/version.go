package utils

import (
	"os/exec"
)

// Returns the version of the app based on the latest tag
// and a possible error
func GetVersion() (string, error) {
	arg0 := "git"
	arg1 := "describe"
	arg2 := "--abbrev=0"
	cmd := exec.Command(arg0, arg1, arg2)

	result, err := cmd.Output()

	return string(result), err
}
