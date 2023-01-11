package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// db is an instance of a Gorm db
var db *gorm.DB

// Site defines the structure of the db
//
// Name: name of the site
// UserName: username for that site
// Password: password for the site | TODO Should probably find a way to encrypt this
type Site struct {
    ID uint64 `json:"id" gorm:"primaryKey`
    Name string `json:"name" gorm:"not null"`
    UserName string `json:"user_name", gorm:"not null"`
    Password string `json:"password", gorm:"not null"`
}

// SetupDb creates a connection to the db
// and initializes the table and columns
func SetupDb(){
    var err error

    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config)
    if err != nil {
        log.Fatal("Could not open db", err)
        panic(err)
    }

    err = db.AutoMigrate(&Site{})
    if err != nil {
        log.Fatal("Could not setup tables and stuff:", err)
    }
}
