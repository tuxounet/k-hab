package utils

import (
	"bytes"
	"text/template"
)

func UnTemplate(ctx *ScopeContext, tpl string, data any) string {

	return ScopingWithReturn(ctx, "utils", "UnTemplate", func(ctx *ScopeContext) string {

		t1 := template.New(tpl)
		t1, err := t1.Parse(tpl)
		ctx.Must(err)
		var buf bytes.Buffer
		tmpl, _ := template.New("tpl").Parse(tpl)
		err = tmpl.Execute(&buf, data)
		ctx.Must(err)
		return buf.String()
	})
}
