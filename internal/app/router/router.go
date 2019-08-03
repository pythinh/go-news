package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pythinh/go-news/internal/app/home"
	"github.com/pythinh/go-news/internal/app/types"
	"github.com/pythinh/go-news/internal/pkg/middleware"
)

type (
	static struct {
		prefix string
		dir    string
	}
)

// Init all routes
func Init() (http.Handler, error) {
	routes := []types.Route{}
	home.NewRouter(&routes)

	r := mux.NewRouter()
	r.Use(middleware.Logging)
	for _, rt := range routes {
		r.HandleFunc(rt.Path, rt.Handler).Methods(rt.Method)
	}

	s := static{"/static/", "web/static"}
	r.PathPrefix(s.prefix).Handler(http.StripPrefix(s.prefix, http.FileServer(http.Dir(s.dir))))
	return r, nil
}
