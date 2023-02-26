package cmd

import (
	"github.com/musaubrian/tinygo/internal/model"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Searches for a specified site records",
	Long:    `Searches for a specified site records(sitename, username and password)`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		model.SearchSite()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
