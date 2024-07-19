package models

import "time"

type Genre struct {
	ID        int
	Genre     string
	Checked   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
