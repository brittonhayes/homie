package templates

import (
	"bytes"
	_ "embed"
	"html/template"

	"github.com/Masterminds/sprig"
)

var (
	//go:embed listing.tmpl
	Listings string
	//go:embed contacted.tmpl
	Contacted string
)

func Render(tpl string, data interface{}) (*bytes.Buffer, error) {
	t, err := template.New("homie").Funcs(sprig.FuncMap()).Parse(tpl)
	if err != nil {
		return nil, err
	}

	b, err := exec(t, data)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func exec(t *template.Template, data interface{}) (*bytes.Buffer, error) {
	var buff bytes.Buffer
	if err := t.Execute(&buff, data); err != nil {
		return nil, err
	}

	return &buff, nil
}
