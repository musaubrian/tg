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
	fmt.Printf("%s: ", prompt)

	result, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Could not read input: ", err)
	}

	result = strings.TrimSuffix(result, "\n")
	return result
}

// Adds a new site's details to the db
func AddSite() {
	var newSite Site
	newSite.Name = getInput("Site Name")
	newSite.UserName = getInput("Site's Username")
	newSite.Password = getInput("Site's Password")
	db.Create(&newSite)
	fmt.Printf("\nSuccessfully added {%s}\n", newSite.Name)
}

// Updates the contents of a specified site
func UpdateSite() {
	var site Site

	siteName := getInput("Site to update")
	db.Where("name = ?", siteName).First(&site)
	fmt.Println("\nEditing: ", site.Name)
	fmt.Printf("Old details\nSiteName: {%s}  UserName {%s}  Password {%s}\n", site.Name, site.UserName, site.Password)
	site.Name = getInput("New site name")
	site.UserName = getInput("New userName")
	site.Password = getInput("New Password")
	db.Save(&site)
	fmt.Println("\nUpdated Successfully")
}

// Delete records associated with a site
func DeleteSite() {
    var site Site
    siteName := getInput("Site to delete")
    db.Where("name = ?", siteName).Delete(&site)
    fmt.Printf("\nDeleted {%s} successfully\n", siteName)
}

// SearchSite searches for a siteName that matches the users input
func SearchSite() {
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
func ListAll() {
	var sites []Site
	db.Find(&sites)
	for _, site := range sites {
		fmt.Println("\nSiteName: ", site.Name)
		fmt.Println("UserName: ", site.UserName)
		fmt.Println("Password: ", site.Password)
	}
}

// Start of user interaction
func TinyGo() {
	fmt.Println("\n// Add[a] | Search[s] | Update[u] | Delete[d] | List all[l]\n")
	startingPoint := getInput("What shall it be?")
	switch startingPoint {
	case "a":
		AddSite()
	case "s":
		SearchSite()
	case "d":
        DeleteSite()
	case "u":
		UpdateSite()
	case "l":
		ListAll()
	default:
		fmt.Println("that option doesn't exist", startingPoint)
	}
}
