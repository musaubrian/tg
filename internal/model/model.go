// Package model sets up the database
// and defines functions that interact with the db
package model

import (
	"path"

	"github.com/musaubrian/tinygo/internal/utils"
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
func SetupDB() error {
	var err error

	homePath, err := utils.GetPath()
	if err != nil {
		return err
	}
	fullPath := path.Join(homePath, "tinygo.db")

	db, err = gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Site{})
	if err != nil {
		return err
	}
	return err
}
