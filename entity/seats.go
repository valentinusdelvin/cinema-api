package entity

import "github.com/google/uuid"

type Seats struct {
	ID          string `gorm:"primary_key"`
	StudioID    uuid.UUID
	SeatsRow    string
	SeatsNumber int
	IsAvailable bool
	Studio      Studio `gorm:"foreignkey:StudioID"`
}
