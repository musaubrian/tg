package model

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

var (
	bold    = color.New(color.Bold)
	success = color.New(color.Bold, color.FgGreen)
	del     = color.New(color.Bold, color.FgRed)
	italic  = color.New(color.Italic, color.Bold)
)

// Get input from the user
func GetInput(prompt string) string {
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
	newSite.Name = GetInput("Site's name")
	newSite.UserName = GetInput("Username")
	newSite.Password = GetInput("Password")
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
	sitename := GetInput("Site to Update")
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

	site.Name = GetInput("New site name")
	if site.Name == "" {
		site.Name = prev_site.Name
		italic.Printf("Reusing previous sitename: %s\n", site.Name)
	}
	site.UserName = GetInput("New userName")

	if site.UserName == "" {
		site.UserName = prev_site.UserName
		italic.Printf("Reusing previous username: %s\n", site.UserName)
	}
	site.Password = GetInput("New Password")
	if site.Password == "" {
		site.Password = prev_site.Password
		italic.Printf("Reusing previous password: %s\n", site.Password)
	}
	db.Save(&site)
	success.Printf("\nUpdated {%s} to {%s}\n\n", sitename, site.Name)
}

// Delete records associated with a site
func DeleteRecord(value string, recordType string) {
	var site Site
	if recordType == "site" {
		db.Where("name = ?", value).Delete(&site)
	} else {
		db.Where("user_name = ?", value).Delete(&site)
	}
	del.Printf("\nDeleted {%s} successfully\n", value)

}

// Returns records matching the users prompt(the site name)
func SearchRecords(value string, recordType string) ([]Site, error) {
	var searchResults []Site
	var nothingFound error
	var result *gorm.DB

	prepd_value := "%" + value + "%"
	if recordType == "sitename" {
		result = db.Raw("SELECT * FROM `Sites` WHERE name LIKE ?", prepd_value).Find(&searchResults)
	} else {
		result = db.Raw("SELECT * FROM `Sites` WHERE user_name LIKE ?", prepd_value).Find(&searchResults)
	}
	if result.RowsAffected == 0 {
		err := fmt.Sprintf("No record found matching [%s]\n", value)
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
