package article

import (
	"net/http"

	"github.com/pythinh/go-news/internal/pkg/types"
)

const (
	get  = http.MethodGet
	post = http.MethodPost
)

// NewRouter append new router
func NewRouter(r *[]types.Route) {
	routes := []types.Route{
		{
			Path:    "/article",
			Method:  get,
			Handler: newRouter().indexHandler,
		},
	}

	*r = append(*r, routes...)
}
