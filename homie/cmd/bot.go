package cmd

import (
	"context"
	"log"

	"github.com/brittonhayes/homie/pkg/bot"
	"github.com/brittonhayes/homie/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// botCmd represents the bot command
var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Start the telegram bot",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancelFunc := context.WithCancel(context.Background())
		defer cancelFunc()

		commands := []bot.Command{
			bot.NewCommand("address", bot.Address),
			bot.NewCommand("status", bot.Status),
			bot.NewCommand("hi", bot.Hi),
			bot.NewCommand("help", bot.Help),
		}

		go func() {
			err := bot.RunWithAllowlist(ctx, viper.GetString("telegram.token"), commands, viper.GetStringSlice("telegram.allowed"))
			if err != nil {
				log.Fatalln(err)
			}
		}()

		logrus.Info("Bot started")
		<-ctx.Done()

		logrus.Info("Bot shutting down")
	},
}

func init() {
	rootCmd.AddCommand(botCmd)
	_, err := config.LoadConfig(".")
	if err != nil {
		return
	}
}
