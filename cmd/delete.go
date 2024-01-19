package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/musaubrian/tg/internal/model"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Remove a record from the database",
	Aliases: []string{"del", "d"},
	// Long: ``,
	Example: `
tinygo delete site
tinygo delete user
	`,
}

var deleteSiteCmd = &cobra.Command{
	Use:   "site",
	Short: "Remove a record using its site name",
	Run: func(cmd *cobra.Command, args []string) {
		siteName := model.GetInput("SiteName")
		model.DeleteRecord(siteName, model.SiteName)
	},
}
var deleteSiteByUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Remove a record using its username",
	Long: `
Remove a record using its username

CAUTION!!
If multiple sites have the same username, they will all be deleted
	`,
	Run: func(cmd *cobra.Command, args []string) {
		userName := model.GetInput("UserName")
		fmt.Printf("This will delete all records with the username [%s]\n", userName)
		agree := model.GetInput("Continue y/N")

		if len(agree) < 1 || agree == "n" {
			fmt.Println("Stopping process")
			os.Exit(1)
		} else if agree == "y" {
			model.DeleteRecord(userName, model.Username)
		} else {
			log.Fatal("Unknown option")
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteSiteCmd)
	deleteCmd.AddCommand(deleteSiteByUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
