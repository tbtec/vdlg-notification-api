package gateway

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/external"
)

type CustomerGateway struct {
	customerService external.ICustomerService
}

func NewCustomerGateway(customerService external.ICustomerService) *CustomerGateway {
	return &CustomerGateway{
		customerService: customerService,
	}
}

func (gtw *CustomerGateway) FindOne(ctx context.Context, id string) (*entity.Customer, error) {
	customer := entity.Customer{}

	slog.InfoContext(ctx, "Finding customer", slog.String("id", id))
	customerResponse, err := gtw.customerService.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	slog.InfoContext(ctx, "Customer response", slog.Any("customerResponse", customerResponse))

	if customerResponse != nil {
		customer = entity.Customer{
			ID:             customerResponse.Content.CustomerId,
			Name:           customerResponse.Content.Name,
			DocumentNumber: customerResponse.Content.DocumentNumber,
			Email:          customerResponse.Content.Email,
			CreatedAt:      customerResponse.Content.CreatedAt,
			UpdatedAt:      customerResponse.Content.UpdatedAt,
		}
		return &customer, nil
	}

	return nil, nil
}
