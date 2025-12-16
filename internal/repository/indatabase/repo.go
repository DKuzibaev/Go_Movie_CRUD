package indatabase

import (
	"go_crud/internal/models"

	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Movie{})
}

func (r *MovieRepository) GetAll() ([]*models.Movie, error) {
	var movies []*models.Movie
	if err := r.db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *MovieRepository) GetByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	if err := r.db.First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *MovieRepository) Update(movie *models.Movie) error {
	return r.db.Save(movie).Error
}

func (r *MovieRepository) Delete(id uint) error {
	return r.db.Delete(&models.Movie{}, id).Error
}
