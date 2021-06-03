package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brittonhayes/homie/pkg/config"
	"github.com/brittonhayes/homie/pkg/parse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List out homes from sheet",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.LoadConfig(".")
		if err != nil {
			log.Fatalln(err)
		}

		p, err := parse.New(c)
		if err != nil {
			log.Fatalln(err)
		}

		// Fetch listings from the sheet
		listings := p.Listings(c.Google.Sheet.HeaderRow)
		b, err := json.MarshalIndent(listings, "", "\t")
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().String("sheet-title", "", "the name of the the sheet tab")
	listCmd.Flags().Int("header-row", 5, "header row containing column headings")
	listCmd.Flags().StringSlice("allowed", []string{}, "list of allowed users")
	_ = viper.BindPFlag("header-row", listCmd.Flags().Lookup("header-row"))
	_ = viper.BindPFlag("sheet-title", listCmd.Flags().Lookup("sheet-title"))
	_ = viper.BindPFlag("allowed", listCmd.Flags().Lookup("allowed"))
}
