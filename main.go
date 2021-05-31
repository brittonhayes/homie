package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brittonhayes/homie/pkg/parse"
	"github.com/brittonhayes/homie/pkg/setup"
)

func main() {
	s, err := setup.Client("", "Listings")
	if err != nil {
		log.Fatalln(err)
	}

	listings := parse.Listings(s)
	bytes, err := json.MarshalIndent(listings, "", "\t")
	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}
