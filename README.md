# Homie 🏡

> A telegram bot and CLI to help me hunt for houses

## Installation

```shell
# Install with go get
go get github.com/brittonhayes/homie/homie
```

## Usage

```shell
## Start the bot
homie bot --config ./.homie.yaml
```

## Config

```yaml
# .homie.yaml
telegram:
  token: "12345"
  allowed:
    - MyUserName

google:
  secrets: "client_secret.json"
  sheet:
    id: "12345"
    header_row: 5
    title: "Listings"

```

### Development

> Add more commands by appending the commands array.

```go

// All commands are of type `CommandFunc func(*tg.Message, config.Configuration) string`

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
```
