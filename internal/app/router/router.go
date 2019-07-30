package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pythinh/go-news/internal/app/home"
	"github.com/pythinh/go-news/internal/app/types"
)

type (
	route struct {
		path    string
		method  string
		handler http.HandlerFunc
	}
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
	for _, rt := range routes {
		r.HandleFunc(rt.Path, rt.Handler).Methods(rt.Method)
	}

	s := static{"/static/", "web/static"}
	r.PathPrefix(s.prefix).Handler(http.StripPrefix(s.prefix, http.FileServer(http.Dir(s.dir))))
	return r, nil
}
