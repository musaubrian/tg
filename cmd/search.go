package cmd

import (
	"fmt"
	"log"

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
		results, err := model.SearchSite()

		if err != nil {
			log.Fatal(err)
		}
		pretty, err := rootCmd.Flags().GetBool("pretty")

		if err != nil {
			log.Fatal(err)

		}
		if pretty {
			t.AddHeader("#", "USER_NAME", "SITE_NAME", "PASSWORD")
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
