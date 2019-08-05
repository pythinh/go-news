package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pythinh/go-news/internal/pkg/respond"
	"github.com/pythinh/go-news/internal/pkg/types"
)

type (
	service interface {
		Get(ctx context.Context, id string) (*types.Article, error)
		GetAll(ctx context.Context) ([]types.Article, error)
		Create(ctx context.Context, article types.Article) (string, error)
		Update(ctx context.Context, article types.Article) error
		Delete(ctx context.Context, id string) error
	}
	// Handler is web handler
	Handler struct {
		srv service
	}
)

// New return new rest api
func New(s service) *Handler {
	return &Handler{s}
}

// Get handle get
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	article, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, article)
}

// GetAll handle get all
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	article, err := h.srv.GetAll(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, article)
}

// Create handle insert
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var article types.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	id, err := h.srv.Create(r.Context(), article)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	article.ID = id
	respond.JSON(w, http.StatusCreated, map[string]string{"id": article.ID})
}

// Update handle modify
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var article types.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	err = h.srv.Update(r.Context(), article)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"id": article.ID})
}

// Delete handle delete
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.srv.Delete(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]string{"status": "success"})
}
