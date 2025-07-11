package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/notification"
)

type NotificationGateway struct {
	notificationService notification.INotificationService
}

func NewNotificationateway(notificationService notification.INotificationService) *NotificationGateway {
	return &NotificationGateway{
		notificationService: notificationService,
	}
}

func (gtw *NotificationGateway) Send(ctx context.Context, notification entity.Notification) error {
	notificationDto := dto.SendNotification{
		Email:   notification.Email,
		Message: notification.Message,
		VideoId: notification.VideoId,
	}

	return gtw.notificationService.Send(ctx, notificationDto)

}
