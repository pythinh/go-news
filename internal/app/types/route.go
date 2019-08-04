package types

import (
	"net/http"

	"github.com/pythinh/go-news/internal/pkg/db"
)

type (
	// DBConns connections to database
	DBConns struct {
		Database db.Connections
	}
	// Route types
	Route struct {
		Path        string
		Method      string
		Handler     http.HandlerFunc
		Middlewares []func(http.HandlerFunc) http.HandlerFunc
	}
)
