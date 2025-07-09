package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type NotificationCreateController struct {
	usc *usecase.NotificationCreateUseCase
}

func NewVideoUpdateController(container *container.Container) *NotificationCreateController {
	return &NotificationCreateController{
		usc: usecase.NewVideoUpdateUseCase(
			gateway.NewNotificationateway(container.NotificationService),
			gateway.NewCustomerGateway(container.CustomerService),
		),
	}
}

func (ctl *NotificationCreateController) Execute(ctx context.Context, notification dto.CreateNotification) error {
	err := ctl.usc.Execute(ctx, notification)
	if err != nil {
		return err
	}
	return nil
}
