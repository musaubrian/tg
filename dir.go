package main

import (
	"os"

	"github.com/musaubrian/tinygo/model"
)

// Creates the 'db/' directory
// Skippes it if the directory exist
func CreateDir() error {
    fullPath := model.GetPath()

    _, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
