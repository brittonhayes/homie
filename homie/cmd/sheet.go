package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sheetCmd represents the sheet command
var sheetCmd = &cobra.Command{
	Use:   "sheet",
	Short: "Get info about the google sheet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nGoogle Sheet Details\n---")
		alignf("ID:", viper.GetString("google.sheet.id"))
		alignf("LINK:", viper.GetString("google.sheet.link"))
		alignf("TITLE: ", viper.GetString("google.sheet.title"))
	},
}

func init() {
	rootCmd.AddCommand(sheetCmd)
}

func alignf(key string, value string) {
	fmt.Printf("%-12s%s\n", key, value)
}
