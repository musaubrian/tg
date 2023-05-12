package cmd

import (
	"log"
	"os"

	"github.com/musaubrian/tinygo/internal/model"
	"github.com/musaubrian/tinygo/internal/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "tinygo",
	Short:   "A cli tool to help manage passwords",
	Long:    `A cli tool to help manage your logins(username and passwords)`,
	Version: "0.4.5",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Create db directory
	if err := utils.CreateDir(); err != nil {
		log.Fatal("Cannot create directory: ", err)
	}
	// Call SetupDb
	if err := model.SetupDB(); err != nil {
		log.Fatal("Db setup error ", err)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tinygo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
