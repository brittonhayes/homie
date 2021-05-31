package parse

import (
	"github.com/Iwark/spreadsheet"
)

const (
	Address int = iota
	City
	Bed
	Bath
	SqFt
	Pet
	Rent
	RelativeToBudget
	Status
	Notes
	Link
	MapLink
)

type Listing struct {
	Address          string `json:"address"`
	City             string `json:"city"`
	Bed              string `json:"bed"`
	Bath             string `json:"bath"`
	SqFt             string `json:"sq_ft"`
	Pet              string `json:"pet"`
	Rent             string `json:"rent"`
	RelativeToBudget string `json:"relative_to_budget"`
	Status           string `json:"status"`
	Notes            string `json:"notes"`
	Link             string `json:"link"`
	MapLink          string `json:"map_link"`
}

func cell(row []spreadsheet.Cell, cell int) string {
	return row[cell].Value
}

func Listings(s *spreadsheet.Sheet) []Listing {
	var listings []Listing
	for i, row := range s.Rows {
		if i > 4 && cell(row, 0) != "" {
			listings = append(listings,
				Listing{
					Address:          cell(row, Address),
					City:             cell(row, City),
					Bed:              cell(row, Bed),
					Bath:             cell(row, Bath),
					SqFt:             cell(row, SqFt),
					Pet:              cell(row, SqFt),
					Rent:             cell(row, Rent),
					RelativeToBudget: cell(row, RelativeToBudget),
					Status:           cell(row, Status),
					Notes:            cell(row, Notes),
					Link:             cell(row, Link),
					MapLink:          cell(row, MapLink),
				})
		}
	}

	return listings
}
