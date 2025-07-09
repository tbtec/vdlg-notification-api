package eventserver

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/env"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/event"
)

type EventServer struct {
	ConsumerService              event.IConsumerService
	NotificationCreateController *controller.NotificationCreateController
}

func NewEventServer(container *container.Container, config env.Config) *EventServer {
	slog.InfoContext(context.Background(), "Creating Event Server...")

	cpc := controller.NewVideoUpdateController(container)
	cs := container.ConsumerService

	return &EventServer{
		ConsumerService:              cs,
		NotificationCreateController: cpc}

}

func (eventServer *EventServer) ConsumeOutput(ctx context.Context) {

	message, err := eventServer.ConsumerService.ConsumeMessageOutput(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error reading message", slog.Any("error", err))
	}
	if message == nil {
		// slog.InfoContext(ctx, "No messages available")
	} else {
		slog.InfoContext(ctx, "Received message: ", &message)

		err2 := eventServer.NotificationCreateController.Execute(ctx, dto.CreateNotification{
			OutputMessage: message,
		})
		if err2 != nil {
			slog.ErrorContext(ctx, "Error processing message", slog.Any("error", err2))

		}
	}
}
