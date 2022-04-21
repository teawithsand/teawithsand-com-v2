package tpl

import "text/template"

const rawTemplate = `<!DOCTYPE html>
<html lang="en">
<head><meta charset="utf-8"><title>{{ .Title }}</title>{{ range .Inputs }}{{ .Render}} {{ end }}{{range .Metas }}{{ .Render}} {{end}}{{range .Links}}{{ .Render }}{{end}}{{range .Scripts}}{{ .Render }}{{end}}</head>
<body>
</body>
</html>`

func getCompiledTemplate() *template.Template {
	tpl, err := template.New("asdf").Parse(rawTemplate)
	if err != nil {
		panic(err)
	}

	return tpl
}

var compiledTemplate = getCompiledTemplate()
