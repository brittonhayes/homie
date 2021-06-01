package bot

import (
	"context"

	"github.com/brittonhayes/homie/pkg/config"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type Commander interface {
	Exec(received *tg.Message, reply *tg.Message) error
}

type Command struct {
	Name   string
	Prefix string
	Fn     CommandFunc
}

type CommandFunc func(*tg.Message, config.Configuration) string

func NewCommand(name string, fn CommandFunc) Command {
	return Command{Name: name, Prefix: name, Fn: fn}
}

func Route(b *tg.BotAPI, update tg.Update, commands []Command) error {
	for _, c := range commands {
		if c.Prefix == update.Message.Command() {
			return c.Exec(b, update.Message)
		}
	}
	return nil
}

func (c *Command) Exec(b *tg.BotAPI, received *tg.Message) error {
	logrus.Infof("BOT: Received command: %s", c.Prefix)

	// Load config
	conf, _ := config.LoadConfig(".")

	// Setup reply
	reply := tg.NewMessage(received.Chat.ID, "")
	reply.ReplyToMessageID = received.MessageID

	// Reply to message
	reply.Text = c.Fn(received, conf)
	_, err := b.Send(reply)
	if err != nil {
		return err
	}

	return nil
}

func RunWithAllowlist(ctx context.Context, token string, commands []Command, allowlist []string) error {
	b, err := tg.NewBotAPI(token)
	if err != nil {
		return err
	}

	var allowed bool
	u := tg.NewUpdate(0)
	u.Timeout = 60
	updates, _ := b.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		for _, user := range allowlist {
			if user == update.Message.From.String() {
				allowed = true
			}
		}

		if !allowed {
			logrus.Errorf("User '%s' not in allow list", update.Message.From.String())
			continue
		}

		err := Route(b, update, commands)
		if err != nil {
			return err
		}
	}

	<-ctx.Done()
	return nil
}

func Run(ctx context.Context, token string, commands []Command) error {
	b, err := tg.NewBotAPI(token)
	if err != nil {
		return err
	}

	u := tg.NewUpdate(0)
	u.Timeout = 60
	updates, _ := b.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		err := Route(b, update, commands)
		if err != nil {
			return err
		}
	}

	<-ctx.Done()
	return nil
}
