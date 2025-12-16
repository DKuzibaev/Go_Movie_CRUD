package handlers

import (
	"encoding/json"
	"go_crud/internal/models"
	"go_crud/internal/repository/indatabase"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type MovieHandler struct {
	store indatabase.MovieStorage
}

func NewMovieHandler(store indatabase.MovieStorage) *MovieHandler {
	return &MovieHandler{store: store}
}

// GetMovies возвращает все фильмы
func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movies, err := h.store.GetAll()
	if err != nil {
		http.Error(w, "failed to get movies", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movies)
}

// CreateMovie создаёт новый фильм
func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	movie.ID = uuid.New().String()

	if err := h.store.Create(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&movie)
}

// GetMovie возвращает фильм по id
func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := mux.Vars(r)["id"]
	idUint64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	id := uint(idUint64)

	movie, err := h.store.GetByID(id)
	if err != nil {
		http.Error(w, "movie not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

// UpdateMovie обновляет фильм по id
func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := mux.Vars(r)["id"]
	idUint64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	id := uint(idUint64)

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	movie.ID = strconv.Itoa(int(id)) // корректное присвоение ID

	if err := h.store.Update(&movie); err != nil { // передаем указатель
		http.Error(w, "movie not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&movie)
}

// DeleteMovie удаляет фильм по id
func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := mux.Vars(r)["id"]
	idUint64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	id := uint(idUint64)

	if err := h.store.Delete(id); err != nil {
		http.Error(w, "movie not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}
