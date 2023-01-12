package model

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// db is an instance of a Gorm db
//var db *gorm.DB

// Site defines the structure of the db
type Site struct {
    ID       uint64 `gorm:"primaryKey not null`
    Name     string `gorm:"not null"`
    UserName string `gorm:"not null"`
    Password string `gorm:"not null"`
}

// SetupDb creates a connection to the db
// and initializes the table and columns
func SetupDb() {
	var err error

	db, err := gorm.Open(sqlite.Open("./db/test.db"), &gorm.Config{})
	if err != nil {
        log.Fatal("Could not open db: ", err)
		panic(err)
	}

	err = db.AutoMigrate(&Site{})
	if err != nil {
		log.Fatal("Could not setup tables and stuff: ", err)
	}
    fmt.Println("Setup database succesfully")
}
