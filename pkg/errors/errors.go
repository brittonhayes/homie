package errors

import (
	"emperror.dev/errors"
	"github.com/sirupsen/logrus"
)

const (
	ErrMsgBot   = "Bot command parsing error"
	ErrStackBot = errors.Sentinel("Bot command parsing error")
)

// func withBotStack() error {
// 	return errors.WithStack(ErrStackBot)
// }

func WithBot(err error) error {
	logrus.Error(err)
	return errors.WithMessage(err, ErrMsgBot)
}
