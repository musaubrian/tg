package model

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	bold    = color.New(color.Bold)
	success = color.New(color.Bold, color.FgGreen)
	del     = color.New(color.Bold, color.FgRed)
	italic  = color.New(color.Italic, color.Bold)
)

// Get input from the user
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

	bold.Println("\n// Adding record")
	newSite.Name = getInput("Site's name")
	newSite.UserName = getInput("Username")
	newSite.Password = getInput("Password")
	db.Create(&newSite)
	success.Printf("\nSuccessfully added {%s}\n", newSite.Name)
}

// Updates the contents of a specified site
func UpdateSite() {
	var (
		site      Site
		prev_site Site
	)

	bold.Println("\n// Updating record")
	sitename := getInput("Site to Update")
	err := db.Where("name = ?", sitename).First(&site)
	if err.Error != nil {
		del.Printf("Record not found")
		return
	}

	prev_site = site
	italic.Println("Press Enter/Return without any value to retain previous values")
	bold.Printf(
		"\nOld details\nSiteName: {%s} UserName: {%s}  Password: {%s}\n",
		site.Name, site.UserName, site.Password)

	site.Name = getInput("New site name")
	if site.Name == "" {
		site.Name = prev_site.Name
		italic.Printf("Reusing previous sitename: %s\n", site.Name)
	}
	site.UserName = getInput("New userName")

	if site.UserName == "" {
		site.UserName = prev_site.UserName
		italic.Printf("Reusing previous username: %s\n", site.UserName)
	}
	site.Password = getInput("New Password")
	if site.Password == "" {
		site.Password = prev_site.Password
		italic.Printf("Reusing previous password: %s\n", site.Password)
	}
	db.Save(&site)
	success.Printf("\nUpdated {%s} to {%s}\n\n", sitename, site.Name)
}

// Delete records associated with a site
func DeleteSite() {
	var site Site
	bold.Println("\n// Deleting record")
	sitename := getInput("Site to delete")
	db.Where("name = ?", sitename).Delete(&site)
	del.Printf("\nDeleted {%s} successfully\n", sitename)
}

// Returns records matching the users prompt(the site name)
func SearchSite() ([]Site, error) {
	var searchResults []Site
	var nothingFound error

	bold.Println("\n// Searching for record")
	sitename := getInput("Site to search for")
	if sitename == "" {
		log.Fatal(del.Println("You need to enter a value"))
	}
	bold.Println("\n//Searching for", sitename)
	prepd_sitename := "%" + sitename + "%"
	result := db.Raw("SELECT * FROM `Sites` WHERE name LIKE ?", prepd_sitename).Find(&searchResults)
	if result.RowsAffected == 0 {
		err := fmt.Sprintf("No site found matching [%s]\n", sitename)
		nothingFound = errors.New(err)
		return searchResults, nothingFound
	}
	return searchResults, nothingFound
}

// Returns all the records in the db
func ListAll() []Site {
	var sites []Site
	db.Find(&sites)
	bold.Println("\n// Listing all records")

	return sites
}
