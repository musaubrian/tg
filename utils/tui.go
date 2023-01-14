package utils

import (
	"github.com/rivo/tview"
)

func Tui() {
	app = tview.NewApplication()
	list := tview.NewList().
		AddItem("Add", "", 'a', func() {
			AddSiteOption()
		}).
		AddItem("Search", "", 's', func() {
			SearchSiteOption()
		}).
		AddItem("Update", "", 'u', func() {
			UpdateSiteOption()
		}).
		AddItem("Delete", "", 'd', func() {
			DeleteSiteOption()
		}).
		AddItem("List all", "", 'l', func() {
			ListAllOption()
		})
	list.SetMainTextColor(tview.Styles.PrimaryTextColor).
		SetSelectedTextColor(tview.Styles.PrimaryTextColor).
		SetSelectedBackgroundColor(tview.Styles.ContrastBackgroundColor)
	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}

