package bot

import (
	"bytes"
	"fmt"

	"github.com/brittonhayes/homie/pkg/config"
	"github.com/brittonhayes/homie/pkg/parse"
	"github.com/brittonhayes/homie/pkg/setup"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

const (
	MsgHousePresent       = "\nLooks like something similar is already in the google sheet! ✅\n\n%s"
	MsgHouseNotPresent    = "Hmmm... couldn't find that one in the sheet 🤔"
	ErrCommandUnknown     = "I don't know that command. Sorry ☹️"
	ErrSomethingWentWrong = "Uh oh. Looks like something went wrong in my wiring!"
)

func Status(received *tg.Message, c config.Configuration) string {
	logrus.Info(received.Chat.ID)
	return "I'm doin' well! 🎉"
}

func Hi(received *tg.Message, c config.Configuration) string {
	return fmt.Sprintf("Hey %s 👋", received.From.FirstName)
}

func Help(received *tg.Message, c config.Configuration) string {
	return "I understand /hi, /status, and /address"
}

func Address(received *tg.Message, c config.Configuration) string {
	s, err := setup.Client(c.Google.Secrets, c.Google.Sheet.ID, c.Google.Sheet.Title)
	if err != nil {
		logrus.Error("failed to setup client", err)
		return ErrSomethingWentWrong
	}

	listings := parse.Listings(s, c.Google.Sheet.HeaderRow)
	similar := parse.SimilarListings(listings, received.CommandArguments(), 0.5)
	logrus.Infof("Parsed listings and found %d similar results", len(similar))
	if len(similar) == 0 {
		return MsgHouseNotPresent
	}

	t, err := parse.Template()
	if err != nil {
		logrus.Error("failed to execute template", err)
		return ErrSomethingWentWrong
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, similar); err != nil {
		logrus.Error("failed to execute template", err)
		return ErrSomethingWentWrong
	}

	return tpl.String()
}
