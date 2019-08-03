package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging log request information
func Logging(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("request path: %s, method: %s, time spent: %s", r.URL.Path, r.Method, time.Since(start))
	})
}
