package router

import (
	"net/http"

	"github.com/pythinh/go-news/internal/pkg/tmpl"
)

type (
	val map[string]interface{}
	// Handler is web handler
	Handler struct{}
)

func homeNew() *Handler {
	return &Handler{}
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request) {
	vals := val{"Title": "Homepage"}
	tmpl.ExecuteTemplate(w, "index.html", vals)
}
