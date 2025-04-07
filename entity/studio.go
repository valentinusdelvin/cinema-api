package entity

import "github.com/google/uuid"

type Studio struct {
	ID       uuid.UUID `gorm:"primary_key"`
	CinemaID uuid.UUID
	Name     string
	Capacity int
	Cinema   Cinema `gorm:"foreignkey:CinemaID"`
	Seats    []Seats
}
