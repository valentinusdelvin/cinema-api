package entity

type Showtime struct {
	ID        int
	MovieID   int
	CinemaID  int
	TheatreID int
	StartTime string
	Time      string
	Movie     Movie  `gorm:"foreignkey:MovieID"`
	Cinema    Cinema `gorm:"foreignkey:CinemaID"`
}
