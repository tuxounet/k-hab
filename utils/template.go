package utils

import (
	"bytes"
	"text/template"
)

func UnTemplate(tpl string, data any) (string, error) {

	var buf bytes.Buffer
	tmpl, err := template.New(tpl).Parse(tpl)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil

}
