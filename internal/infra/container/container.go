package container

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/tbtec/tremligeiro/internal/env"
	"github.com/tbtec/tremligeiro/internal/infra/event"
	"github.com/tbtec/tremligeiro/internal/infra/external"
	"github.com/tbtec/tremligeiro/internal/infra/notification"
)

type Container struct {
	Config              env.Config
	ConsumerService     event.IConsumerService
	NotificationService notification.INotificationService
	CustomerService     external.ICustomerService
}

func New(config env.Config) (*Container, error) {
	factory := Container{}
	factory.Config = config

	return &factory, nil
}

func (container *Container) Start(ctx context.Context) error {

	var awsConfig aws.Config
	var err error

	if container.Config.Env == "local-stack" { // LocalStack
		awsConfig = container.GetLocalStackConfig(ctx)
	} else {
		awsConfig, err = config.LoadDefaultConfig(ctx,
			config.WithRegion(container.Config.AwsRegion))
		log.Printf("AWS Region: %s", container.Config.AwsRegion)
		if err != nil {
			log.Fatalf("erro ao carregar config: %v", err)
		}
	}

	container.ConsumerService = event.NewConsumerService(container.Config.OutputQueueUrl, awsConfig)
	container.NotificationService = notification.NewNotificationService(container.Config.SMTPUser, container.Config.SMTPPass)
	container.CustomerService = external.NewCustomerService(getCustomerConf(container.Config))

	return nil
}

func (container *Container) Stop() error {
	return nil
}

func getCustomerConf(config env.Config) external.CustomerConfig {
	return external.CustomerConfig{
		Url: config.CustomerUrl,
	}
}

func (container *Container) GetLocalStackConfig(ctx context.Context) aws.Config {

	awsConfig, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
	)
	awsConfig.BaseEndpoint = aws.String("http://localhost:4566")

	if err != nil {
		log.Fatalf("erro ao carregar config: %v", err)
	}

	return awsConfig
}
