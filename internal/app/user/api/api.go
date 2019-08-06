package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pythinh/go-news/internal/pkg/login"
	"github.com/pythinh/go-news/internal/pkg/respond"
	"github.com/pythinh/go-news/internal/pkg/types"
)

type (
	service interface {
		Get(ctx context.Context, username string) (*types.User, error)
		Check(ctx context.Context, username string) bool
		GetAll(ctx context.Context) ([]types.User, error)
		Create(ctx context.Context, user *types.User) (string, error)
		CreatePass(ctx context.Context, username, password string) error
		Update(ctx context.Context, user *types.User) error
		UpdatePass(ctx context.Context, username, oldPassword, newPassword string) error
		Delete(ctx context.Context, id string) error
	}
	// Handler is web handler
	Handler struct {
		srv           service
		localLoginSrv login.Service
	}
)

// New return new rest api
func New(s service, localLogin login.Service) *Handler {
	return &Handler{s, localLogin}
}

// Get handle get
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	user, err := h.srv.Get(r.Context(), username)
	user.Password = ""
	if err != nil {
		respond.Error(w, err, http.StatusNotFound)
	}
	respond.JSON(w, http.StatusOK, user)
}

// GetAll handle get all
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	user, err := h.srv.GetAll(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, user)
}

// Create handle create
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	user := types.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	exits := h.srv.Check(r.Context(), user.Username)
	if exits {
		respond.JSON(w, http.StatusConflict, map[string]string{"status": "409", "message": "user already existed"})
		return
	}
	userName, err := h.srv.Create(r.Context(), &user)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, userName)
}

// CreatePass handle create pass
func (h *Handler) CreatePass(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	user := types.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	err = h.srv.CreatePass(r.Context(), username, user.Password)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, username)
}

// Update handle update
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	user := types.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	err = h.srv.Update(r.Context(), &user)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, user.Username)
}

// UpdatePass handle update pass
func (h *Handler) UpdatePass(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	user := types.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	err = h.srv.UpdatePass(r.Context(), username, user.OldPassword, user.Password)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, username)
}

// Delete handle delete
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.srv.Delete(r.Context(), id)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, id)
}
