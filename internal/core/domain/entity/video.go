package entity

import (
	"time"

	"github.com/tbtec/tremligeiro/internal/types/ulid"
)

type Video struct {
	ID             string
	CustomerId     string
	Status         VideoStatus
	FileNameInput  string
	FileNameOutput string
	UploadUrl      *string
	DownloadUrl    *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type VideoStatus string

const (
	VideoStatusReceived   VideoStatus = "RECEIVED"
	VideoStatusProcessing VideoStatus = "PROCESSING"
	VideoStatusCompleted  VideoStatus = "COMPLETED"
	VideoStatusError      VideoStatus = "ERROR"
)

func NewVideo(customerId string) Video {
	return Video{
		ID:             ulid.NewUlid().String(),
		CustomerId:     customerId,
		Status:         VideoStatusReceived,
		FileNameInput:  "",
		FileNameOutput: "",
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
}

func (video *Video) SetFileNameInput(fileName string) {
	video.FileNameInput = fileName
}

func (video *Video) SetFileNameOutput(fileName string) {
	video.FileNameOutput = fileName
}

func (video *Video) SetUploadUrl(uploadUrl string) {
	video.UploadUrl = &uploadUrl
}

func (video *Video) SetDownloadUrl(downloadUrl string) {
	video.DownloadUrl = &downloadUrl
}

func (video *Video) SetStatus(status VideoStatus) {
	video.Status = status
}
