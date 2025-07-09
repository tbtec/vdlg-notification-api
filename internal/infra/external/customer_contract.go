package external

import "time"

type CustomerConfig struct {
	Url string
}

type CustomerContent struct {
	Content CustomerResponse `json:"content"`
}
type CustomerResponse struct {
	CustomerId     string    `json:"id"`
	Name           string    `json:"name"`
	DocumentNumber string    `json:"documentNumber"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
