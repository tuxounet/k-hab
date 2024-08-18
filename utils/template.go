package utils

import (
	"bytes"
	"text/template"
)

func UnTemplate(ctx *ScopeContext, tpl string, data any) string {

	return ScopingWithReturn(ctx, "utils", "UnTemplate", func(ctx *ScopeContext) string {

		var buf bytes.Buffer
		tmpl, _ := template.New(tpl).Parse(tpl)
		err := tmpl.Execute(&buf, data)
		ctx.Must(err)
		return buf.String()
	})
}
