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
	tmpl.ExecuteTemplate(w, "index.html", vals)
}
