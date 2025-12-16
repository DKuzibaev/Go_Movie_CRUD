package server

import (
	"go_crud/internal/handlers"
	"go_crud/internal/repository/indatabase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	store  indatabase.MovieStorage
	router *mux.Router
}

func NewServer(store indatabase.MovieStorage) *Server {
	s := &Server{
		store:  store,
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	handler := handlers.NewMovieHandler(s.store)

	s.router.HandleFunc("/movies", handler.GetMovies).Methods("GET")
	s.router.HandleFunc("/movies/{id}", handler.GetMovie).Methods("GET")
	s.router.HandleFunc("/movies", handler.CreateMovie).Methods("POST")
	s.router.HandleFunc("/movies/{id}", handler.UpdateMovie).Methods("PUT")
	s.router.HandleFunc("/movies/{id}", handler.DeleteMovie).Methods("DELETE")
}

func (s *Server) Start() {
	log.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", s.router))
}
