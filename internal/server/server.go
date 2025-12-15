package server

import (
	"go_crud/internal/handlers"
	"go_crud/internal/repository/inmemory"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	store  *inmemory.MovieStore
	router *mux.Router
}

func NewServer(store *inmemory.MovieStore) *Server {
	s := &Server{
		store:  store,
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	r := s.router
	handler := handlers.NewMovieHandler(s.store)

	r.HandleFunc("/movies", handler.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", handler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handler.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handler.DeleteMovie).Methods("DELETE")

}

func (s *Server) Start() {
	log.Println("Starting server at port 8000")

	log.Fatal(http.ListenAndServe(":8000", s.router))
}
