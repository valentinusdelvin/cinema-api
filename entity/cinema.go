package entity

import "github.com/google/uuid"

type Cinema struct {
	ID          uuid.UUID `gorm:"primary_key"`
	Name        string
	Address     string
	TotalStudio int
	Studio      []Studio
}
