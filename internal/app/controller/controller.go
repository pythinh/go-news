package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pythinh/go-news/internal/app/article"
	"github.com/pythinh/go-news/internal/app/home"
	"github.com/pythinh/go-news/internal/pkg/db"
	"github.com/pythinh/go-news/internal/pkg/db/mongodb"
	"github.com/pythinh/go-news/internal/pkg/middleware"
	"github.com/pythinh/go-news/internal/pkg/types"
)

type static struct {
	prefix string
	dir    string
}

// DBConns connections to database
type DBConns struct {
	Database db.Connections
}

// InitRoute all routes
func InitRoute(conns *DBConns) (http.Handler, error) {
	routes := []types.Route{}
	home.NewRouter(&routes)
	article.NewRouter(&routes, &conns.Database)

	r := mux.NewRouter()
	r.Use(middleware.Logging)
	for _, rt := range routes {
		r.HandleFunc(rt.Path, rt.Handler).Methods(rt.Method)
	}

	s := static{"/static/", "web/static"}
	r.PathPrefix(s.prefix).Handler(http.StripPrefix(s.prefix, http.FileServer(http.Dir(s.dir))))
	return r, nil
}

// InitDB connections to database
func InitDB(conf *types.Server) *DBConns {
	conns := &DBConns{}
	conns.Database.Type = conf.DB.Type

	switch conf.DB.Type {
	case db.TypeMongoDB:
		s, err := mongodb.Dial(&conf.DB.ConfigDB)
		if err != nil {
			log.Panicln("failed to dial to target server, err:", err)
		}
		conns.Database.MongoDB = s
	}
	return conns
}

// Close all underlying connections
func (c *DBConns) Close() {
	c.Database.Close()
}
