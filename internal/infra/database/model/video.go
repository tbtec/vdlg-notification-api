package model

import "time"

type Video struct {
	ID             string    `gorm:"column:video_id;primaryKey"`
	CustomerId     string    `gorm:"column:customer_id"`
	Status         string    `gorm:"column:status"`
	FileNameInput  string    `gorm:"column:file_name_input"`
	FileNameOutput string    `gorm:"column:file_name_output"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}
