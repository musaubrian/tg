//Package model sets up the database
// and defines functions that interact with the db
package model

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Get input from the user
func getInput(prompt string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("%s: ",prompt)

    result, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal("Could not read input: ", err)
    }

    result = strings.TrimSuffix(result, "\n")
    return result
}

// adds a new site's details to the db
func addSite() {
    var newSite Site
    newSite.Name = getInput("Site Name")
    newSite.UserName = getInput("Site's Username")
    newSite.Password = getInput("Site's Password")
    db.Create(&newSite)
}

func updateSite(){}
func DeleteSite(){}
func SearchSite(){}
func listAll(){
    var sites []Site
    result := db.Find(&sites)
    fmt.Println(result)
}


 // Start of user interaction
 func TinyGo(){
    fmt.Println("\n// Add[a] | Search[s] | Update[u] | Delete[d]\n")
    startingPoint := getInput("What shall it be?")
    switch startingPoint {
    case "a":
        addSite()
    case "s":
        fmt.Println("chose s")
    case "d":
        fmt.Println("chose d")
    case "u":
        fmt.Println("chose u")
    default:
        fmt.Println("that option doesn't exist", startingPoint)
    }
}
