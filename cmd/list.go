package cmd

import (
	"fmt"
	"log"

	"github.com/musaubrian/tg/internal/model"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all records in the database",
	Long: `Lists all the records(sitename, username, password) available 
    in the database`,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		pretty, err := rootCmd.Flags().GetBool("pretty")
		if err != nil {
			log.Fatal(err)
		}
		sites := model.ListAll()
		if pretty {
			t.AddHeader("\n#", "USERNAME", "SITE_NAME", "PASSWORD")
			for i, site := range sites {
				t.AddLine(i+1, site.UserName, site.Name, site.Password)
			}
			t.Print()
		} else {
			for _, site := range sites {
				fmt.Println("\nSiteName:", site.Name)
				fmt.Println("UserName:", site.UserName)
				fmt.Println("Password:", site.Password)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
