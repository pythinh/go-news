package article

import (
	"log"
	"net/http"

	"github.com/pythinh/go-news/internal/app/article/api"
	"github.com/pythinh/go-news/internal/pkg/db"
	"github.com/pythinh/go-news/internal/pkg/types"
)

const (
	get  = http.MethodGet
	post = http.MethodPost
)

// NewRouter append new router
func NewRouter(r *[]types.Route, conns *db.Connections) {
	var repo repository
	switch conns.Type {
	case db.TypeMongoDB:
		repo = newMongoRepository(conns.MongoDB)
	default:
		log.Panicln("database type not supported:", conns.Type)
	}
	srv := newService(repo)
	routes := []types.Route{
		// Route
		{
			Path:    "/article",
			Method:  get,
			Handler: newRouter().indexHandler,
		},
		// API
		{
			Path:    "/api/article",
			Method:  get,
			Handler: api.New(srv).GetAll,
		},
	}

	*r = append(*r, routes...)
}
