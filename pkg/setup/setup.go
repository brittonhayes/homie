package setup

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/Iwark/spreadsheet"
	"golang.org/x/oauth2/google"
)

const (
	ErrReadFile            = "failed to read secrets file: %s"
	ErrParsingJWT          = "failed to parse JWT from JSON: %s"
	ErrFetchingSpreadsheet = "failed to fetch spreadsheet from id: %s"
)

func Client(secretFilePath string, sheetID string, sheetTitle string) (*spreadsheet.Sheet, error) {
	data, err := ioutil.ReadFile(secretFilePath)
	if err != nil {
		return nil, fmt.Errorf(ErrReadFile, err)
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return nil, fmt.Errorf(ErrParsingJWT, err)
	}

	client := conf.Client(context.TODO())
	svc := spreadsheet.NewServiceWithClient(client)
	sheet, err := svc.FetchSpreadsheet(sheetID)
	if err != nil {
		return nil, fmt.Errorf(ErrFetchingSpreadsheet, err)
	}

	s, err := sheet.SheetByTitle(sheetTitle)
	if err != nil {
		return nil, err
	}

	return s, nil
}
