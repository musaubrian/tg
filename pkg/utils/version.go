package utils

import (
	"log"
	"os/exec"
)

// Returns the version of the app based on the latest tag
func GetVersion() string{
    arg0 := "git"
    arg1 := "describe"
    arg2 := "--abbrev=0"
    cmd := exec.Command(arg0, arg1, arg2)

    result, err := cmd.Output()
    if err != nil {
        log.Fatal("Could not get the version from tags")
    }
    return string(result)
}
