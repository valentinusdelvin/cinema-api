package usecase

import (
	"Cinema_API/internal/repository"
	"Cinema_API/pkg/bcrypt"
	"Cinema_API/pkg/jwt"
)

type Usecase struct {
	UserUsecase  UserUsecase
	MovieUsecase MovieUsecase
}

type InitializersParam struct {
	Repository repository.Repository
	Bcrypt     bcrypt.BcryptService
	JWT        jwt.JwtService
}

func UsecaseInit(param InitializersParam) *Usecase {
	UserUsecase := UserUsecaseInit(param.Repository.UserRepository, param.Bcrypt, param.JWT)
	MovieUsecase := MovieUsecaseInit(param.Repository.MovieRepository)

	return &Usecase{
		UserUsecase:  UserUsecase,
		MovieUsecase: MovieUsecase,
	}
}
