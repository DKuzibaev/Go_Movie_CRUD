package main

import (
	"go_crud/internal/repository/indatabase"
	"go_crud/internal/server"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=127.0.0.1 user=go_user password=go_password dbname=go_crud port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	repo := indatabase.NewMovieRepository(db)

	// АВТОМИГРАЦИЯ
	if err := repo.Migrate(); err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(repo)
	srv.Start()
}
