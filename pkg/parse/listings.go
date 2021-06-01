package parse

import (
	"github.com/Iwark/spreadsheet"
	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

const (
	Address int = iota
	City
	Bed
	Bath
	SqFt
	Pets
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
	Pets             string `json:"pets"`
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

func SimilarListings(listings []Listing, comparison string, similarityLevel float64) []Listing {
	var similar []Listing
	for _, listing := range listings {
		if strutil.Similarity(listing.Address, comparison, metrics.NewHamming()) >= similarityLevel {
			similar = append(similar, listing)
		}
	}

	return similar
}

func Listings(s *spreadsheet.Sheet, header int) []Listing {
	var listings []Listing
	for i, row := range s.Rows {
		if i > header && cell(row, Address) != "" {
			listings = append(listings,
				Listing{
					Address:          cell(row, Address),
					City:             cell(row, City),
					Bed:              cell(row, Bed),
					Bath:             cell(row, Bath),
					SqFt:             cell(row, SqFt),
					Pets:             cell(row, Pets),
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
