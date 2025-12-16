package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID         string    `gorm:"primaryKey" json:"id"`
	Isbn       string    `json:"isbn"`
	Title      string    `json:"title"`
	DirectorID uint      `json:"director_id"` // внешняя колонка
	Director   *Director `json:"director" gorm:"foreignKey:DirectorID"`
	CreatedAt  time.Time `json:"created_at"`
}

type Director struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
