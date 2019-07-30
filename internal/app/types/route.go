package types

import "net/http"

// Route types
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}
