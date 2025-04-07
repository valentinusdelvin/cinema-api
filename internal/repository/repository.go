package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository  UserRepository
	MovieRepository MovieRepository
}

func RepositoryInit(db *gorm.DB) *Repository {
	UserRepository := UserRepositoryInit(db)
	MovieRepository := MovieRepositoryInit(db)

	return &Repository{
		UserRepository:  UserRepository,
		MovieRepository: MovieRepository,
	}
}
