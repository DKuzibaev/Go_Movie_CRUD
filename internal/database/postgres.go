package database

import (
	"log"

	"go_crud/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=go_user password=go_password dbname=go_crud port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// 🔥 Автомиграция
	if err := db.AutoMigrate(&models.Movie{}, &models.Director{}); err != nil {
		log.Fatal("migration failed:", err)
	}

	log.Println("Database connected & migrated")
	return db
}
