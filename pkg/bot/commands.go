package bot

import (
	"fmt"
	"os"
	"time"

	"github.com/brittonhayes/homie/pkg/config"
	"github.com/brittonhayes/homie/pkg/parse"
	"github.com/brittonhayes/homie/pkg/templates"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

const (
	MsgHouseNotPresent    = "Hmmm... couldn't find that one in the sheet 🤔"
	ErrSomethingWentWrong = "Uh oh. Looks like something went wrong in my wiring!"
	ErrFailedTemplate     = "failed to execute template"
)

func Status(received *tg.Message, c config.Configuration) string {
	logrus.Info(received.Chat.ID)
	return "I'm doin' well! 🎉"
}

func Hi(received *tg.Message, _ config.Configuration) string {
	return fmt.Sprintf("Hey %s 👋", received.From.FirstName)
}

func Help(*tg.Message, config.Configuration) string {
	return "I understand /hi, /status, and /address"
}

func Address(received *tg.Message, c config.Configuration) string {
	p, err := parse.New(c)
	if err != nil {
		logrus.Error("failed to setup client", err)
		return ErrSomethingWentWrong
	}

	// Fetch listings from the sheet
	listings := p.Listings(c.Google.Sheet.HeaderRow)

	// Look for listings similar to the user's message
	similar := p.SimilarListings(listings, received.CommandArguments(), 0.5)

	logrus.Infof("Parsed listings and found %d similar results", len(similar))
	if len(similar) == 0 {
		return MsgHouseNotPresent
	}

	// Create the listing template
	b, err := templates.Render(templates.Listings, similar)
	if err != nil {
		logrus.Error(ErrFailedTemplate, err)
		return ErrSomethingWentWrong
	}

	return b.String()
}

func Goodnight(received *tg.Message, c config.Configuration) string {
	go func() {
		time.Sleep(5 * time.Second)
		logrus.Infof("Going to sleep now!")
		os.Exit(1)
	}()

	return "Have a good one! I'm clocking out for the evening. \nhttps://media.tenor.com/images/df51877535a3e38c9cccd2f23ff154a2/tenor.gif"
}

func Contacted(received *tg.Message, c config.Configuration) string {
	p, err := parse.New(c)
	if err != nil {
		logrus.Error("failed to setup client", err)
		return ErrSomethingWentWrong
	}

	// Fetch listings from the sheet
	listings := p.Listings(c.Google.Sheet.HeaderRow)

	var contacted []parse.Listing
	for _, l := range listings {
		if l.Status == "Contacted" || l.Status == "Applied" {
			contacted = append(contacted, l)
		}
	}

	// Create the listing template
	b, err := templates.Render(templates.Contacted, contacted)
	if err != nil {
		logrus.Error(ErrFailedTemplate, err)
		return ErrSomethingWentWrong
	}

	return b.String()
}
