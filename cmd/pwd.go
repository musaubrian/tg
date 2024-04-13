package cmd

import (
	"fmt"
	"log"

	"github.com/musaubrian/tg/internal/utils"
	"github.com/spf13/cobra"
)

// pwdCmd represents the pwd command
var pwdCmd = &cobra.Command{
	Use:     "pwd",
	Short:   "Generate a 15 character string as password",
	Aliases: []string{"p"},
	Long: `
Generate a 15 character string as password
parse 'c' as argument to copy it without displaying it
	`,
	Run: func(cmd *cobra.Command, args []string) {
		pwd := utils.GeneratePassword()

		copyPwd, err := rootCmd.Flags().GetBool("copy")
		if err != nil {
			log.Fatal(err)
		}
		if copyPwd {
			utils.CopyToClipboard(pwd)
			fmt.Println("Copied to clipboard")
			return
		}
		fmt.Println(pwd)
	},
}

func init() {
	rootCmd.AddCommand(pwdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pwdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pwdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
