package repository

import (
	"Cinema_API/entity"
	"Cinema_API/models"
	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovie(movie models.CreateMovie) error
	GetAllMovies() ([]models.MovieKey, error)
}

func MovieRepositoryInit(db *gorm.DB) MovieRepository {
	return &movieRepository{
		db: db,
	}
}

type movieRepository struct {
	db *gorm.DB
}

func (m *movieRepository) CreateMovie(movie models.CreateMovie) error {
	err := m.db.Create(&movie).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *movieRepository) GetAllMovies() ([]models.MovieKey, error) {
	var movies []models.MovieKey
	err := m.db.Model(entity.Movie{}).Find(&movies).Error
	if err != nil {
		return nil, err
	}
	return movies, nil
}
