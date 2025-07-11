package entity

type Notification struct {
	Email   string
	Message string
	VideoId string
}

func NewNotification(email, message, video string) Notification {
	return Notification{
		Email:   email,
		Message: message,
		VideoId: video,
	}
}
