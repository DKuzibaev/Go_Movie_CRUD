package main

import (
	"fmt"
	"go_crud/internal/repository/inmemory"
	"go_crud/internal/server"
)

func main() {
	store := inmemory.NewMovieStore()

	fmt.Println(store.GetAll())

	s := server.NewServer(store)
	s.Start()
}
