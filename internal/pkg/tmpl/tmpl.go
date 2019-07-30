package tmpl

import (
	"html/template"
	"io"
	"log"
)

// ExecuteTemplate applies the template
func ExecuteTemplate(w io.Writer, name string, data interface{}) {
	var tmpls = template.Must(template.ParseGlob("web/template/*.html"))
	err := tmpls.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Panic(err)
	}
}
