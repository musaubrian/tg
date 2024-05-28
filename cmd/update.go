package cmd

import (
	"github.com/musaubrian/tg/internal/model"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a single site records",
	Long: `
Update a single site records either by the user_name or site_name
NOTE:
  When there are records with the same user_name, only the first one gets updated`,
	Aliases: []string{"u"},
}

var updatebyUserName = &cobra.Command{
	Use:   "user",
	Short: "Update a record by the username",
	Run: func(cmd *cobra.Command, args []string) {
		model.UpdateRecord(model.Username)
	},
}
var updatebySiteName = &cobra.Command{
	Use:   "site",
	Short: "Update a record by the sitename",
	Run: func(cmd *cobra.Command, args []string) {
		model.UpdateRecord(model.SiteName)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updatebySiteName)
	updateCmd.AddCommand(updatebyUserName)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
