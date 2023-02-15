package cmd

import (
	"fmt"
	"log"

	"github.com/musaubrian/tinygo/pkg/utils"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Displays tinygo's version",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		version, err := utils.GetVersion()
        if err != nil {
            log.Fatal("Could not get version from git tag")
        }
		fmt.Println("tinygo", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
