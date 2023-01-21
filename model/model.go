// Package model sets up the database
// and defines functions that interact with the db
package model

import (
	"log"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// db defines a gorm instance
// https://gorm.io
var db *gorm.DB

// Site defines the structure of the db
type Site struct {
	ID      int32
	Name      string
	UserName string
	Password string
}

// SetupDb creates a connection to the db
// and initializes the table and columns
func SetupDb() {
	var err error

	homePath := GetPath()
	fullPath := path.Join(homePath, "tinygo.db")

	db, err = gorm.Open(sqlite.Open(fullPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not open db: ", err)
		panic(err)
	}

	err = db.AutoMigrate(&Site{})
	if err != nil {
		log.Fatal("Could not setup tables and stuff: ", err)
	}
}
