package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/musaubrian/tinygo/model"
)

// TODO:: raise error when no flag is parsed
// Set flag for when to only generate passwords
// or just us the db
func CheckFlag() {

	useGenerator := flag.Bool("p", false, "Generate password")
	useDb := flag.Bool("d", false, "Access records")
	flag.Parse()

	if *useGenerator {
		text := GeneratePassword()
		fmt.Printf("\n%s\n", text)
	}
	if *useDb {
		if err := CreateDir(); err != nil {
			log.Fatal("Could not create directory", err)
		}
		model.SetupDb()
		model.TinyGo()
	}
}
