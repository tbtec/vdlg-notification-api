package notification

import (
	"context"
	"log"
	"log/slog"
	"net/smtp"

	"github.com/tbtec/tremligeiro/internal/dto"
)

type INotificationService interface {
	Send(ctx context.Context, notification dto.SendNotification) error
}

type NotificationService struct {
	Email    string
	Password string
}

func NewNotificationService(email, password string) *NotificationService {
	return &NotificationService{
		Email:    email,
		Password: password,
	}
}

func (service *NotificationService) Send(ctx context.Context, notification dto.SendNotification) error {

	from := service.Email
	password := service.Password

	to := []string{notification.Email}
	subject := "[Video Ligeiro] Processamento de vídeo concluído"
	body := notification.Message

	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}

	slog.InfoContext(ctx, "Email sent successfully")

	return nil

}
