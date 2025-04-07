package models

import "github.com/google/uuid"

type UserRegister struct {
	ID          uuid.UUID
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UserLogin struct {
	Email    string
	Password string
}

type UserLoginResponse struct {
	Token string
}
