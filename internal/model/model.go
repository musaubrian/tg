// Package model sets up the database
// and defines functions that interact with the db
package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Site defines the structure of the db
type Site struct {
	ID       int32
	Name     string
	UserName string
	Password string
}

// db defines a gorm instance
// https://gorm.io
var db *gorm.DB

// SetupDb creates a connection to the db
// and initializes the table and columns
// Returns a possible error
func SetupDB(dbLoc string) error {
	var err error

	db, err = gorm.Open(sqlite.Open(dbLoc), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Site{})
	if err != nil {
		return err
	}
	return err
}
