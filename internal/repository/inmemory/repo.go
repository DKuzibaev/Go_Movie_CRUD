package inmemory

import "go_crud/internal/models"

type MovieStore struct {
	movies []models.Movie
}

func NewMovieStore() *MovieStore {
	return &MovieStore{
		movies: []models.Movie{},
	}
}

func (s *MovieStore) GetAll() []models.Movie {
	return s.movies
}
