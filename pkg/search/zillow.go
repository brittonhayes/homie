package search

import (
	"github.com/jmank88/zillow"
)

func Zillow(token string) zillow.Zillow {
	return zillow.New(token)
}
