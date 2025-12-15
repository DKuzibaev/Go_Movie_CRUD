package handlers

import (
	"encoding/json"
	"net/http"

	"go_crud/internal/repository/inmemory"
)

type MovieHandler struct {
	store *inmemory.MovieStore
}

func NewMovieHandler(store *inmemory.MovieStore) *MovieHandler {
	return &MovieHandler{store: store}
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies := h.store.GetAll()
	json.NewEncoder(w).Encode(movies)
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {

}
