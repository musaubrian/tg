package main

import "os"

// Creates the 'db/' directory
// Skippes it if the directory exist
func CreateDir() error {
	_, err := os.Stat("db")
	if os.IsNotExist(err) {
		err := os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
