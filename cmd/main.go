package main

import (
	"go_crud/internal/repository/inmemory"
	"go_crud/internal/server"
)

func main() {
	store := inmemory.NewMovieStore()
	s := server.NewServer(store)
	s.Start()
}
