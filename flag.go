package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/musaubrian/tinygo/model"
)

// Set flag for when to only generate passwords
// use the db or display the help message
func CheckFlag() {

	useGenerator := flag.Bool("p", false, "Generate password")
	useDb := flag.Bool("d", false, "Access records")
    help := flag.Bool("h", false, "Display usage instructions")
	flag.Parse()

    switch {
    case *useGenerator:
        text := GeneratePassword()
        fmt.Printf("\n%s\n", text)
    case *useDb:
        if err := CreateDir(); err != nil {
            log.Fatal("Could not create directory", err)}


        model.SetupDb()
        model.TinyGo()
    case *help:
        usageInstructions()
    default:
        usageInstructions()
    }
}

// Print usage instructions
func usageInstructions()  {
    fmt.Println(`
Usage of tinygo:
    -d  Access records
    -h  Display usage instructions
    -p  Generate passwords
    `)
}
