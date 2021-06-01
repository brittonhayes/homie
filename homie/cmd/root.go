package cmd

import (
	"fmt"
	"os"

	"github.com/brittonhayes/homie/pkg/config"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "homie",
	Short: "Help you find houses",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	c, err := config.LoadConfig(".")
	if err != nil {
		return
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&c.File, "config", "", "config file (default is $HOME/.homie.yaml)")
	rootCmd.PersistentFlags().StringVar(&c.Google.Sheet.ID, "sheet", "", "the google sheet id")
	rootCmd.PersistentFlags().StringVar(&c.Google.Secrets, "secrets", "", "the path to the google service account secrets file")
	rootCmd.PersistentFlags().StringVarP(&c.Telegram.Token, "telegram-token", "t", "", "the telegram api token")

	_ = viper.BindPFlag("sheet", rootCmd.PersistentFlags().Lookup("sheet"))
	_ = viper.BindPFlag("secrets", rootCmd.PersistentFlags().Lookup("secrets"))
	_ = viper.BindPFlag("telegram-token", rootCmd.PersistentFlags().Lookup("telegram-token"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	c, err := config.LoadConfig(".")
	if err != nil {
		return
	}

	if c.File != "" {
		viper.SetConfigFile(c.File)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".homie")
	}
}
