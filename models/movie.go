package models

import (
	"github.com/google/uuid"
	"mime/multipart"
)

type CreateMovie struct {
	ID         uuid.UUID             `form:"id"`
	Title      string                `form:"title"`
	Synopsis   string                `form:"synopsis"`
	Year       int                   `form:"year"`
	Rating     float64               `form:"rating"`
	Genre      string                `form:"genre"`
	Duration   int                   `form:"duration"`
	Director   string                `form:"director"`
	Actors     []string              `form:"actors"`
	PosterLink string                `form:"poster_link"`
	PosterIMG  *multipart.FileHeader `form:"poster_img"`
}

type MovieKey struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Synopsis   string    `json:"synopsis"`
	PosterLink string    `json:"poster_link"`
	Rating     float64   `json:"rating"`
}
