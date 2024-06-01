package models

import "time"

type Movie struct {
	ID          int
	Title       string
	ReleaseDate time.Time
	Runtime     int
	MppaRating  string
	Description string
}
