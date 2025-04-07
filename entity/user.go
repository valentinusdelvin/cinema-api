package entity

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `gorm:"primary_key"`
	Name         string
	Email        string
	Password     string
	PhoneNumber  string
	IsAdmin      bool
	IsSubscribed bool
}
