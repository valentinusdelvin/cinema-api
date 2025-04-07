package entity

import "github.com/google/uuid"

type Movie struct {
	ID       uuid.UUID `gorm:"primary_key"`
	Title    string
	Synopsis string
	Year     int
	Rating   float64
	Genre    string
	Duration int
	Director string
	Actors   []string
	Poster   string
}
