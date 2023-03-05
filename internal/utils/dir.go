// Package utils defines utility functions
package utils

import (
	"log"
	"os"
	"path"
)

// Get full path to homeDir
func GetPath() (string, error) {
	var fullPath string
	var err error

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return fullPath, err
	}
	fullPath = path.Join(homeDir, ".db")

	return fullPath, err
}

// Creates the 'db/' directory
// Skips it if the directory exist
func CreateDir() error {
	var err error

	fullPath, err := GetPath()
	if err != nil {
		return err
	}

	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
