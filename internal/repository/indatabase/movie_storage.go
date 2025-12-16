package indatabase

import "go_crud/internal/models"

type MovieStorage interface {
	GetAll() ([]*models.Movie, error)
	GetByID(id uint) (*models.Movie, error)
	Create(movie *models.Movie) error
	Update(movie *models.Movie) error
	Delete(id uint) error
}
