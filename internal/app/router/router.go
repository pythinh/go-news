package router

import (
	"net/http"

	"github.com/gorilla/mux"
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

const (
	get  = http.MethodGet
	post = http.MethodPost
)

// Init all routes
func Init() (http.Handler, error) {
	homeView := homeNew()
	routes := []route{
		// home
		{"/", get, homeView.index},
	}
	r := mux.NewRouter()
	for _, rt := range routes {
		r.HandleFunc(rt.path, rt.handler).Methods(rt.method)
	}

	s := static{"/static/", "web/static"}
	r.PathPrefix(s.prefix).Handler(http.StripPrefix(s.prefix, http.FileServer(http.Dir(s.dir))))
	return r, nil
}
