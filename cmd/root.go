package cmd

import (
	"log"
	"os"
	"path"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/musaubrian/tg/internal/model"
	"github.com/musaubrian/tg/internal/utils"
	"github.com/spf13/cobra"
)

var t = table.New().StyleFunc(func(row, col int) lipgloss.Style {
	return lipgloss.NewStyle().Width(26)
})

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tg",
	Short: "A cli tool to help manage passwords",
	Long:  `A cli tool to help manage your logins(username and passwords)`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Create db directory
	if err := utils.CreateDir(); err != nil {
		log.Fatal("Cannot create directory: ", err)
	}

	homePath, err := utils.GetPath()
	if err != nil {
		log.Fatal(err)
	}
	fullPath := path.Join(homePath, "tinygo.db")

	if err := model.SetupDB(fullPath); err != nil {
		log.Fatal("Db setup error ", err)
	}

	err = rootCmd.Execute()
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
	rootCmd.PersistentFlags().BoolP("pretty", "p", true, "List results in a nice table format")
	rootCmd.PersistentFlags().BoolP("copy", "c", false, "Copy password to clipboard")
}
