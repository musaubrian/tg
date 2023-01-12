// Package model sets up the database
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
    fmt.Printf("Successfully added {%s}\n", newSite.Name)
}

func updateSite(){}
func deleteSite(){}
func searchSite(){
    var site Site
    fmt.Println("//Searching")
    name := getInput("Site name?")
    result := db.Where("name = ?", name).First(&site)
    if result.RowsAffected == 0 {
        fmt.Println("No site found matching", name)
    } else {
        fmt.Println("\nSiteName:", site.Name)
        fmt.Println("Username:", site.UserName)
        fmt.Println("Password:", site.Password)
    }

}

// Lists all the records in the db
func listAll(){
    var sites []Site
    db.Find(&sites)
    for _, site := range sites {
        fmt.Println("\nSiteName: ", site.Name)
        fmt.Println("UserName: ", site.UserName)
        fmt.Println("Password: ", site.Password)
    }    
}


 // Start of user interaction
 func TinyGo(){
    fmt.Println("\n// Add[a] | Search[s] | Update[u] | Delete[d] | List all[l]\n")
    startingPoint := getInput("What shall it be?")
    switch startingPoint {
    case "a":
        addSite()
    case "s":
        searchSite()
    case "d":
        fmt.Println("chose d")
    case "u":
        fmt.Println("chose u")
    case "l":
        listAll()
    default:
        fmt.Println("that option doesn't exist", startingPoint)
    }
}
