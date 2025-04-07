package bcrypt

import "golang.org/x/crypto/bcrypt"

type BcryptService interface {
	GenerateHashPassword(password string) (string, error)
	CompareHashPassword(hashedPassword, password string) error
}
type bcryptService struct {
	Cost int
}

func BcryptInit() BcryptService {
	return &bcryptService{
		Cost: 10,
	}
}

func (b *bcryptService) GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.Cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func (b *bcryptService) CompareHashPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
