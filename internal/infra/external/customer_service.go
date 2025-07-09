package external

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-resty/resty/v2"
	"github.com/tbtec/tremligeiro/internal/infra/httpclient"
)

type ICustomerService interface {
	FindOne(ctx context.Context, id string) (*CustomerContent, error)
}

type CustomerService struct {
	httpclient *resty.Client
	config     CustomerConfig
}

func NewCustomerService(config CustomerConfig) ICustomerService {
	return &CustomerService{
		config:     config,
		httpclient: httpclient.New(),
	}
}

func (service *CustomerService) FindOne(ctx context.Context, id string) (*CustomerContent, error) {
	customerResponse := CustomerContent{}

	url := service.config.Url
	path := "/api/v1/customer" + "/" + id

	slog.InfoContext(ctx, "Finding customer", slog.String("id", id), slog.String("url", url+path))

	response, err := service.httpclient.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&customerResponse).
		Get(url + path)
	if err != nil {
		return nil, err
	}
	slog.InfoContext(ctx, "Response from customer service", slog.Any("response", response))

	if response.StatusCode() != 200 {
		return nil, errors.New("failed to find customer: " + response.Status())
	}

	slog.InfoContext(ctx, "Customer found", slog.Any("customerResponse", customerResponse))
	return &customerResponse, nil
}
