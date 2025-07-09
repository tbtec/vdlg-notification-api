package usecase

import (
	"context"
	"log/slog"
	"strings"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type NotificationCreateUseCase struct {
	notificationGtw *gateway.NotificationGateway
	customerGtw     *gateway.CustomerGateway
}

func NewVideoUpdateUseCase(notificationGtw *gateway.NotificationGateway, customerGtw *gateway.CustomerGateway) *NotificationCreateUseCase {
	return &NotificationCreateUseCase{
		notificationGtw: notificationGtw,
		customerGtw:     customerGtw,
	}
}

func (uc *NotificationCreateUseCase) Execute(ctx context.Context, updateVideo dto.CreateNotification) error {

	slog.InfoContext(ctx, "Executing video update use case", slog.Any("updateVideo", updateVideo))

	if updateVideo.OutputMessage != nil {
		slog.InfoContext(ctx, "Processing OutputMessage", slog.Any("OutputMessage", updateVideo.OutputMessage))

		customerId := uc.getCustomerId(updateVideo.OutputMessage.FileName)
		slog.InfoContext(ctx, "CustomerId", slog.Any("customer", customerId))

		customer, err := uc.customerGtw.FindOne(ctx, customerId)
		if err != nil {
			slog.ErrorContext(ctx, "Error finding customer", slog.Any("error", err))
			return err
		}

		if customer != nil {
			slog.InfoContext(ctx, "Customer found", slog.Any("customer", customer))

			notification := entity.NewNotification(
				customer.Email,
				"Your video processing is complete. You can access it at")

			uc.notificationGtw.Send(ctx, notification)
		}
	}

	return nil
}

func (uc *NotificationCreateUseCase) getCustomerId(fileName string) string {
	parts := strings.Split(fileName, ".")
	file := strings.Split(parts[0], "/")[1]

	return strings.Split(file, "_")[1]
}
