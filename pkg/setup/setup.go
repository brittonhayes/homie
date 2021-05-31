package setup

import (
	"github.com/Iwark/spreadsheet"
)

func Client(sheetID string, sheetTitle string) (*spreadsheet.Sheet,error) {
	svc,err := spreadsheet.NewService()
	if err != nil {
		return nil, err
	}

	sheet, err := svc.FetchSpreadsheet(sheetID)
	if err != nil {
		return nil, err
	}

	s, err := sheet.SheetByTitle(sheetTitle)
	if err != nil {
		return nil, err
	}

	return s,nil
}
