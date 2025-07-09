package entity

import "time"

type Customer struct {
	ID             string
	Name           string
	DocumentNumber string
	Email          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
