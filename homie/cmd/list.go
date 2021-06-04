package cmd

import (
	"fmt"
	"log"

	"github.com/brittonhayes/homie/pkg/config"
	"github.com/brittonhayes/homie/pkg/parse"
	"github.com/brittonhayes/homie/pkg/templates"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"ls"},
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
		listings := p.Listings(viper.GetInt("google.sheet.header_row"))
		output, err := templates.Render(templates.Listings, listings)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
