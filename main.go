package main

import (
	"log"

	"github.com/musaubrian/tinygo/model"
)

func main() {
	if err := CreateDir(); err != nil {
		log.Fatal("Could not create directory", err)
	}

	model.SetupDb()
    model.TinyGo()
}
