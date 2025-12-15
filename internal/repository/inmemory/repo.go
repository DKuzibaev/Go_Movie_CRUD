package inmemory

import (
	"errors"
	"go_crud/internal/models"
	"sync"

	"github.com/google/uuid"
)

type MovieStore struct {
	movies []models.Movie
	mu     sync.Mutex
}

func NewMovieStore() *MovieStore {
	return &MovieStore{
		movies: []models.Movie{},
	}
}

func (s *MovieStore) GetAll() []models.Movie {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]models.Movie, len(s.movies))
	copy(result, s.movies)
	return result
}

func (s *MovieStore) Create(movie models.Movie) models.Movie {
	s.mu.Lock()
	defer s.mu.Unlock()

	movie.ID = uuid.NewString()
	s.movies = append(s.movies, movie)
	return movie
}

func (s *MovieStore) GetByID(id string) (models.Movie, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, item := range s.movies {
		if item.ID == id {
			return item, nil
		}
	}

	return models.Movie{}, errors.New("movie not found")
}

func (s *MovieStore) UpdateByID(id string, newMovie models.Movie) (models.Movie, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, item := range s.movies {
		if item.ID == id {
			newMovie.ID = id       // сохраняем старый ID
			s.movies[i] = newMovie // обновляем элемент в слайсе
			return newMovie, nil
		}
	}

	return models.Movie{}, errors.New("movie not found")
}

func (s *MovieStore) DeleteByID(id string) (models.Movie, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, movie := range s.movies {
		if movie.ID == id {
			// Сохраняем удалённый фильм для ответа
			deleted := movie
			// Удаляем элемент из слайса
			s.movies = append(s.movies[:i], s.movies[i+1:]...)
			return deleted, nil
		}
	}

	return models.Movie{}, errors.New("movie not found")
}
