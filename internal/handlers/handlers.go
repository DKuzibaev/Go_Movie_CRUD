package handlers

import (
	"encoding/json"
	"net/http"

	"go_crud/internal/models"
	"go_crud/internal/repository/inmemory"

	"github.com/gorilla/mux"
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
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	h.store.Create(movie)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	movie, err := h.store.GetByID(id)
	if err != nil {
		http.Error(w, "movie not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	var newMovie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedMovie, err := h.store.UpdateByID(id, newMovie)
	if err != nil {
		http.Error(w, "movie not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedMovie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	deletedMovie, err := h.store.DeleteByID(id)
	if err != nil {
		http.Error(w, "movie not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(deletedMovie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
