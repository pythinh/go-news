package tmpl

import (
	"html/template"
	"io"
	"log"
)

// ExecuteTemplate applies the template
func ExecuteTemplate(w io.Writer, name string, data interface{}) {
	rootDir := "web/template/"
	baseFile := rootDir + "base.html"
	appFile := rootDir + name
	
	var tmpls = template.Must(template.ParseFiles(baseFile, appFile))
	err := tmpls.Execute(w, data)
	if err != nil {
		log.Panicln(err)
	}
}
