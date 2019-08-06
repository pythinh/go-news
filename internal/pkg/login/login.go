package login

import (
	"log"
	"net/http"

	"github.com/pythinh/go-news/internal/pkg/respond"
)

type (
	// Service login
	Service interface {
		Authenticate(username, password string) (interface{}, error)
	}
	// Handle login
	Handle struct {
		localLoginSrv Service
	}
)

// New login handle
func New(localLogin Service) *Handle {
	return &Handle{localLogin}
}

// Authenticate handle
func (h *Handle) Authenticate(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Printf("login with Username: %s, Password: ***", username)
	respondValue, err := h.localLoginSrv.Authenticate(username, password)
	if err != nil {
		respond.Error(w, err, http.StatusForbidden)
		return
	}
	respond.JSON(w, http.StatusAccepted, respondValue)
}
