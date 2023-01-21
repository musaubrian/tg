package model

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// Get full path to homeDir
func GetPath() string {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fullPath := path.Join(homeDir, ".db")

	return fullPath
}

// get input from the user
func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("$ %s: ", prompt)

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

	fmt.Println("\n// Adding record")
	newSite.Name = getInput("Site's name")
	newSite.UserName = getInput("Username")
	newSite.Password = getInput("Password")
	db.Create(&newSite)
	fmt.Printf("\nSuccessfully added {%s}\n", newSite.Name)
}

// Updates the contents of a specified site
func UpdateSite() {
	var site Site

	fmt.Println("\n// Updating record")
	sitename := getInput("Site to Update")
	db.Where("name = ?", sitename).First(&site)

	fmt.Printf("Old details\nSiteName: {%s} UserName: {%s}  Password: {%s}\n", site.Name, site.UserName, site.Password)
	site.Name = getInput("New site name")
	site.UserName = getInput("New userName")
	site.Password = getInput("New Password")
	db.Save(&site)
	fmt.Printf("\nUpdated {%s}", site.Name)
}

// Delete records associated with a site
func DeleteSite() {
	var site Site

	sitename := getInput("Site to delete")
	db.Where("name = ?", sitename).Delete(&site)
	fmt.Printf("\nDeleted {%s} successfully\n", sitename)
}

// SearchSite searches for a siteName that matches the users input
func SearchSite() {
	var site Site

	sitename := getInput("Site to search for")
	fmt.Println("\n//Searching for", sitename)
	result := db.Where("name = ?", sitename).First(&site)
	if result.RowsAffected == 0 {
		fmt.Println("No site found matching", sitename)
	} else {
		fmt.Println("\nSiteName:", site.Name)
		fmt.Println("Username:", site.UserName)
		fmt.Printf("Password: %s\n\n", site.Password)
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
	fmt.Println(`
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                           TinyGo

    Add[a] | Update[u] | Search[s] | Delete[d] | ListAll[l]

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    `)
	choice := getInput("What shall it be?")
	switch choice {
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
		fmt.Printf("Option [%s] doesn't exist\n", choice)
	}
}
