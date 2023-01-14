package utils

import (

	"github.com/musaubrian/tinygo/model"
	"github.com/rivo/tview"
)

func AddSiteOption() {
	form.AddInputField("Enter site name", "", 20, nil, nil).
		SetLabelColor(tview.Styles.PrimaryTextColor).
		SetFieldBackgroundColor(tview.Styles.ContrastBackgroundColor)

	form.AddInputField("Enter username", "", 20, nil, nil).
		SetLabelColor(tview.Styles.PrimaryTextColor).
		SetFieldBackgroundColor(tview.Styles.ContrastBackgroundColor)

	form.AddInputField("Enter password", "", 20, nil, nil).
		SetLabelColor(tview.Styles.PrimaryTextColor).
		SetFieldBackgroundColor(tview.Styles.ContrastBackgroundColor)

	form.AddButton("Submit", func() {
		site := form.GetFormItem(0).(*tview.InputField).GetText()
		user := form.GetFormItem(1).(*tview.InputField).GetText()
		pass := form.GetFormItem(2).(*tview.InputField).GetText()
		if site != "" && user != "" && pass != "" {
			model.AddSite(site, user, pass)
			app.SetRoot(tview.NewTextView().SetText("Site added successfully!"), true)
		} else {
			app.SetRoot(tview.NewTextView().SetText("Please fill all the fields"), true)
		}
	})
}

// SearchSiteOption function to search for a site by name
func SearchSiteOption() {
	form.AddInputField("Enter site name to search", "", 20, nil, nil).
		SetLabelColor(tview.Styles.PrimaryTextColor).
		SetFieldBackgroundColor(tview.Styles.ContrastBackgroundColor)
	form.AddButton("Search", func() {
		sitename := form.GetFormItem(0).(*tview.InputField).GetText()
		if sitename != "" {
			model.SearchSite(sitename)
		} else {
			app.SetRoot(tview.NewTextView().SetText("Please enter a site name"), true)
		}
	})
}

// UpdateSiteOption function to update a site by name
func UpdateSiteOption() {
	form.AddInputField("Enter site name to update", "", 20, nil, nil).
		SetLabelColor(tview.Styles.PrimaryTextColor).
		SetFieldBackgroundColor(tview.Styles.ContrastBackgroundColor)
	form.AddButton("Update", func() {
		sitename := form.GetFormItem(0).(*tview.InputField).GetText()
		if sitename != "" {
			model.UpdateSite(sitename)
		} else {
			app.SetRoot(tview.NewTextView().SetText("Please enter a site name"), true)
		}
	})
}

// DeleteSiteOption function to delete a site by name
func DeleteSiteOption() {
	form.AddInputField("Enter site name to delete", "", 20, nil, nil).
		SetLabelColor(tview.Styles.PrimaryTextColor).
		SetFieldBackgroundColor(tview.Styles.ContrastBackgroundColor)
	form.AddButton("Delete", func() {
		sitename := form.GetFormItem(0).(*tview.InputField).GetText()
		if sitename != "" {
			model.DeleteSite(sitename)
		} else {
			app.SetRoot(tview.NewTextView().SetText("Please enter a site name"), true)
		}
	})
}

// ListAllOption function to list all the sites in the database
func ListAllOption() {
	form.AddButton("List all", func() {
		model.ListAll()
	})
}
