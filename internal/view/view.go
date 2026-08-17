package view

import (
	"embed"
	"html/template"
	"io"

	"github.com/danielronalds/danielronalds.github.io/internal/site"
)

//go:embed templates/root.tmpl.html
var templateFiles embed.FS

var rootTemplate = template.Must(template.ParseFS(templateFiles, "templates/root.tmpl.html"))

type rootData struct {
	Content template.HTML
}

func Render(writer io.Writer, content site.HTMLContent) error {
	return rootTemplate.ExecuteTemplate(writer, "root", rootData{
		Content: template.HTML(content),
	})
}
