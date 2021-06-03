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

Available Commands:
  bot         Start the telegram bot
  help        Help about any command
  list        List out homes from sheet
  sheet       Get info about the google sheet

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

## Development

> Add more commands by appending the commands array.

```go
// All commands are of type `CommandFunc func(*tg.Message, config.Configuration) string`
commands := []bot.Command{
    bot.NewCommand("address", bot.Address),
    bot.NewCommand("status", bot.Status),
    bot.NewCommand("hi", bot.Hi),
    bot.NewCommand("help", bot.Help),
    bot.NewCommand("goodnight", bot.Goodnight)
}
```

```go
// Example commands that will make the bot go to sleep
func Goodnight(received *tg.Message, c config.Configuration) string {
	go func() {
		time.Sleep(5 * time.Second)
		logrus.Infof("Going to sleep now!")
		os.Exit(1)
	}()

	return "Have a good one! I'm clocking out for the evening. \nhttps://media.tenor.com/images/df51877535a3e38c9cccd2f23ff154a2/tenor.gif"
}
```
