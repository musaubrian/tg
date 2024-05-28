package model

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

type RecordType int

const (
	Username RecordType = iota
	SiteName
)

var (
	bold    = color.New(color.Bold)
	success = color.New(color.Bold, color.FgGreen)
	del     = color.New(color.Bold, color.FgRed)
	italic  = color.New(color.Italic, color.Bold)
)

// Get input from the user
func GetInput(prompt string) string {
	var result string

	huh.NewInput().Title(prompt).Value(&result).Run()

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
func UpdateRecord(recordType RecordType) {
	var site Site

	bold.Println("\n// Updating record")

	if recordType == SiteName {
		sitename := GetInput("Site record to Update")
		err := db.Where("name = ?", sitename).First(&site)
		if err.Error != nil {
			del.Printf("Record not found")
			return
		}
		updatedSite := UpdateRec(site)
		printChanges(site, updatedSite)
	} else if recordType == Username {
		user := GetInput("User record to Update")
		err := db.Where("user_name = ?", user).First(&site)
		if err.Error != nil {
			del.Printf("Record not found")
			return
		}
		updatesSite := UpdateRec(site)
		printChanges(site, updatesSite)
	}

}

// Delete records associated with a site
func DeleteRecord(value string, recordType RecordType) {
	var site Site
	if recordType == SiteName {
		db.Where("name = ?", value).Delete(&site)
	} else {
		db.Where("user_name = ?", value).Delete(&site)
	}
	del.Printf("\nDeleted {%s} successfully\n", value)

}

// Returns records matching the users prompt(the site name)
func SearchRecords(value string, searchBy RecordType) ([]Site, error) {
	var searchResults []Site
	var nothingFound error
	var result *gorm.DB

	prepd_value := "%" + value + "%"
	if searchBy == SiteName {
		result = db.Raw("SELECT * FROM `Sites` WHERE name LIKE ?", prepd_value).Find(&searchResults)
	} else if searchBy == Username {
		result = db.Raw("SELECT * FROM `Sites` WHERE user_name LIKE ?", prepd_value).Find(&searchResults)
	} else {
		return searchResults, errors.New("invalid option")
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

func UpdateRec(site Site) Site {
	prev_site := site
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
	return site
}

func printChanges(prev, updated Site) {
	originalValue := reflect.ValueOf(prev)
	updatedValue := reflect.ValueOf(updated)
	for i := 0; i < originalValue.NumField(); i++ {
		origField := originalValue.Field(i)
		updtField := updatedValue.Field(i)
		if !reflect.DeepEqual(origField.Interface(), updtField.Interface()) {
			field := bold.Sprintf("%v", originalValue.Type().Field(i).Name)
			prevVal := bold.Sprintf("%v", origField.Interface())
			newVal := bold.Sprintf("%v", updtField.Interface())
			success.Printf("%s changed from %s to %s\n", field, prevVal, newVal)
		}
	}
}
