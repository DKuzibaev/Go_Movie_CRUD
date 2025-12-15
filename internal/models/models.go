package models

import "time"

type Movie struct {
	ID        string    `json:"id"`
	Isbn      string    `json:"isbn"`
	Title     string    `json:"title"`
	Director  *Director `json:"director"`
	CreatedAt time.Time `json:"created_at"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
