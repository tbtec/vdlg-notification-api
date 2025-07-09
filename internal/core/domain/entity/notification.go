package entity

type Notification struct {
	Email   string
	Message string
}

func NewNotification(email, message string) Notification {
	return Notification{
		Email:   email,
		Message: message,
	}
}
