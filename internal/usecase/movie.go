package usecase

import (
	"Cinema_API/internal/repository"
	"Cinema_API/models"
)

type MovieUsecase interface {
	CreateMovie(movie models.CreateMovie) error
	GetAllMovies() ([]models.MovieKey, error)
}

type movieUsecase struct {
	movieRepo repository.MovieRepository
}

func MovieUsecaseInit(movieRepo repository.MovieRepository) MovieUsecase {
	return &movieUsecase{
		movieRepo: movieRepo,
	}
}

func (m *movieUsecase) CreateMovie(movie models.CreateMovie) error {
	err := m.movieRepo.CreateMovie(movie)
	if err != nil {
		return err
	}
	return nil
}

func (m *movieUsecase) GetAllMovies() ([]models.MovieKey, error) {
	movies, err := m.movieRepo.GetAllMovies()
	if err != nil {
		return nil, err
	}
	return movies, nil
}
