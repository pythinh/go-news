package home

import (
	"net/http"

	"github.com/pythinh/go-news/internal/pkg/tmpl"
)

type (
	val     map[string]interface{}
	handler struct{}
)

func newRouter() *handler {
	return &handler{}
}

func (h *handler) indexHandler(w http.ResponseWriter, r *http.Request) {
	vals := val{"Title": "Homepage"}
	tmpl.ExecuteTemplate(w, "home/index.html", vals)
}

func (h *handler) aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home/about.html", nil)
}
