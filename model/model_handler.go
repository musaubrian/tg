package model

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Get input from the user
func GetInput(prompt string) string {
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
func AddSite(siteName string, username string, password string) {
	var newSite Site
	newSite.Name = siteName
	newSite.UserName = username
	newSite.Password =password
    db.Create(&newSite)
//	fmt.Printf("\nSuccessfully added {%s}\n", newSite.Name)
}

// Updates the contents of a specified site
func UpdateSite(sitename string) {
	var site Site

	db.Where("name = ?", sitename).First(&site)
	fmt.Println("\nEditing: ", site.Name)
	fmt.Printf("Old details\nSiteName: {%s}  UserName {%s}  Password {%s}\n", site.Name, site.UserName, site.Password)
	site.Name = GetInput("New site name")
	site.UserName = GetInput("New userName")
	site.Password = GetInput("New Password")
	db.Save(&site)
	fmt.Println("\nUpdated Successfully")
}

// Delete records associated with a site
func DeleteSite(sitename string) {
	var site Site
	db.Where("name = ?", sitename).Delete(&site)
	fmt.Printf("\nDeleted {%s} successfully\n", sitename)
}

// SearchSite searches for a siteName that matches the users input
func SearchSite(sitename string) {
	var site Site

	fmt.Println("//Searching")
	result := db.Where("name = ?", sitename).First(&site)
	if result.RowsAffected == 0 {
		fmt.Println("No site found matching", sitename)
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
	startingPoint := GetInput("What shall it be?")
	switch startingPoint {
	case "a":
	//	AddSite()
	case "s":
	//	SearchSite()
	case "d":
	//	DeleteSite()
	case "u":
	//	UpdateSite()
	case "l":
	//	ListAll()
	default:
		fmt.Println("that option doesn't exist", startingPoint)
	}
}
