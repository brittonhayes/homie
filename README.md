# Homie 🏡

> A Telegram bot and CLI to help me hunt for houses

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

## Spreadsheet

> Homie expects that you've got a spreadsheet in Google sheets that has these headers.

|Address|City|Bed|Bath|Sq Ft|Pets|Rent|Relative to Budget|Status |Notes |
|:---: |:---: |:---:|:---: |:---: | :---: |:---: |:---: |:---: | :---: |
| 1234 Example street | Example City | 3 | 2 | 1400 | Yes | 2650 | -300 | Contacted | Looks like a nice place |

## Configuration

Configuring homie depends on a your `.homie.yaml` and a `client_secret.json` which you can download from the Google
Console.

<details>
<summary>View .homie.yaml</summary>
<br>

> .homie.yaml

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

</details>

<details>
<summary>View client_secret.json</summary>
<br>

```json
{
  "type": "service_account",
  "project_id": "",
  "private_key_id": "",
  "private_key": "",
  "client_email": "",
  "client_id": "",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "h"
}
```

</details>

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

<details>
<summary>Example command</summary>
<br>

```go
// Example commands that will make the bot go to sleep
func Goodnight(received *tg.Message, c config.Configuration) string {
    go func () {
        time.Sleep(5 * time.Second)
        logrus.Infof("Going to sleep now!")
    os.Exit(1)
    }()

    return "Have a good one! I'm clocking out for the evening. \nhttps://media.tenor.com/images/df51877535a3e38c9cccd2f23ff154a2/tenor.gif"
}
```
</details>

