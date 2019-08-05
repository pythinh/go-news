package respond

import (
	"encoding/json"
	"net/http"
)

// JSON write status, data to http response writer
func JSON(w http.ResponseWriter, status int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

// Error write error, status to http response writer
func Error(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}
