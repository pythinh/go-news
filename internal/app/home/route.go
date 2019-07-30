package home

import (
	"net/http"

	"github.com/pythinh/go-news/internal/app/types"
)

const (
	get = http.MethodGet
)

// NewRouter append new router
func NewRouter(r *[]types.Route) {
	routes := []types.Route{
		{
			Path:    "/",
			Method:  get,
			Handler: newRouter().indexHandler,
		},
	}
	
	*r = append(*r, routes...)
}
