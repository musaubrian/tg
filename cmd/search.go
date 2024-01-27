package cmd

import (
	"fmt"
	"log"

	"github.com/musaubrian/tg/internal/model"
	"github.com/musaubrian/tg/internal/utils"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Searches for a specified site records",
	Long:    `Searches for a specified site records(sitename, username and password)`,
	Aliases: []string{"s"},
}

var searchByUserName = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Search for site by username",
	Example: `

tg search user some_username
tinygo search user some_username -p
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("You did not parse value to search for")
		}
		pretty, err := rootCmd.Flags().GetBool("pretty")
		if err != nil {
			log.Fatal(err)
		}
		results, err := model.SearchRecords(args[0], model.Username)
		if err != nil {
			log.Fatal(err)
		}

		if len(results) == 1 {
			utils.CopyToClipboard(results[0].Password)
			fmt.Printf("Copied [%s's] password for [%s] to clipboard\n", results[0].UserName, results[0].Name)
			return
		}
		if pretty {
			t.AddHeader("\n#", "USER_NAME", "SITE_NAME", "PASSWORD")
			for i, v := range results {
				t.AddLine(i+1, v.UserName, v.Name, v.Password)
			}
			t.Print()
		} else {
			for _, site := range results {
				fmt.Println("\nSiteName:", site.Name)
				fmt.Println("UserName:", site.UserName)
				fmt.Println("Password:", site.Password)
			}
		}

	},
}

var searchbySiteName = &cobra.Command{
	Use:   "site",
	Short: "Search for site by sitename",
	Example: `
tinygo search site some_sitename
tinygo search site some_sitename -p
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("You did not parse value to search for")
		}
		pretty, err := rootCmd.Flags().GetBool("pretty")
		if err != nil {
			log.Fatal(err)
		}
		results, err := model.SearchRecords(args[0], model.SiteName)
		if err != nil {
			log.Fatal(err)
		}
		if len(results) == 1 {
			utils.CopyToClipboard(results[0].Password)
			fmt.Printf("Copied [%s's] password for [%s] to clipboard\n", results[0].UserName, results[0].Name)
			return
		}
		if pretty {
			t.AddHeader("\n#", "USER_NAME", "SITE_NAME", "PASSWORD")
			for i, v := range results {
				t.AddLine(i+1, v.UserName, v.Name, v.Password)
			}
			t.Print()
		} else {
			for _, site := range results {
				fmt.Println("\nSiteName:", site.Name)
				fmt.Println("UserName:", site.UserName)
				fmt.Println("Password:", site.Password)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.AddCommand(searchByUserName)
	searchCmd.AddCommand(searchbySiteName)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
