package usecase

import (
	"Cinema_API/entity"
	"Cinema_API/internal/repository"
	"Cinema_API/models"
	"Cinema_API/pkg/bcrypt"
	"Cinema_API/pkg/jwt"
	"github.com/google/uuid"
)

type UserUsecase interface {
	Register(param models.UserRegister) error
	Login(param models.UserLogin) (models.UserLoginResponse, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
	bcrypt   bcrypt.BcryptService
	jwtAuth  jwt.JwtService
}

func UserUsecaseInit(userRepo repository.UserRepository, bcrypt bcrypt.BcryptService, jwtAuth jwt.JwtService) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		bcrypt:   bcrypt,
		jwtAuth:  jwtAuth,
	}
}
func (u *userUsecase) Register(param models.UserRegister) error {
	hashedPassword, err := u.bcrypt.GenerateHashPassword(param.Password)
	if err != nil {
		return err
	}

	user := entity.User{
		ID:          uuid.New(),
		Name:        param.Name,
		Email:       param.Email,
		Password:    string(hashedPassword),
		PhoneNumber: param.PhoneNumber,
	}

	err = u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Login(param models.UserLogin) (models.UserLoginResponse, error) {
	result := models.UserLoginResponse{}

	user, err := u.userRepo.GetUser(models.UserLogin{
		Email: param.Email,
	})
	if err != nil {
		return result, err
	}

	err = u.bcrypt.CompareHashPassword(user.Password, param.Password)
	if err != nil {
		return result, err
	}

	token, err := u.jwtAuth.GenerateToken(user.ID, user.IsAdmin, user.IsSubscribed)
	if err != nil {
		return result, err
	}

	result.Token = token
	return result, nil
}
