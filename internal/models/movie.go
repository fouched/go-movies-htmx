package models

import "time"

type Movie struct {
	ID          int
	Title       string
	ReleaseDate time.Time
	RunTime     int
	MPAARating  string
	Description string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
