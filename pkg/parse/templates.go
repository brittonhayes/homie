package parse

import (
	"html/template"

	"github.com/Masterminds/sprig"
)

func Template() (*template.Template, error) {
	return template.New("listings").Funcs(sprig.FuncMap()).Parse(`Looks like something similar is already in the google sheet! ✅

{{range .}}
🏡	Address: 	{{.Address}}
🏙️	City:		{{.City}}
💸	Rent:		{{.Rent}}
🛏️	Bed:		{{.Bed}}
🛀	Bath:		{{.Bath}}
📏	SqFt:		{{.SqFt}}
🐕	Pets:		{{.Pets}}
🚦	Status:		{{.Status}}
🤏	Budget:		{{.RelativeToBudget}}
🔗 	Link:		{{.Link}}
🗺️	Map:		https://www.google.com/maps/search/?api=1&query={{ nospace .Address }}

{{end}}
`)
}
