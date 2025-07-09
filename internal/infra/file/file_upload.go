package file

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type IFileService interface {
	GenerateUploadUrl(ctx context.Context, fileName string) (string, error)
	GenerateDownloadUrl(ctx context.Context, fileName string) (string, error)
}

type FileUploadService struct {
	BucketName string
	Client     *s3.Client
}

func NewFileService(bucketName string, config aws.Config) IFileService {
	return &FileUploadService{
		BucketName: bucketName,
		Client:     s3.NewFromConfig(config),
	}
}

func (service *FileUploadService) GenerateUploadUrl(ctx context.Context, fileName string) (string, error) {
	slog.InfoContext(ctx, "Generating upload URL", slog.String("fileName", fileName))

	key := "input/" + fileName
	expiration := 15 * time.Minute

	presignClient := s3.NewPresignClient(service.Client)

	req, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(service.BucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expiration))

	if err != nil {
		slog.InfoContext(ctx, "Erro ao gerar URL pré-assinada: %v", err)
		return "", fmt.Errorf("Erro ao gerar URL pré-assinada: %w", err)
	}

	return req.URL, nil
}

func (service *FileUploadService) GenerateDownloadUrl(ctx context.Context, fileName string) (string, error) {
	slog.InfoContext(ctx, "Generating download URL", slog.String("fileName", fileName))

	key := "output/" + fileName
	expiration := 15 * time.Minute

	presignClient := s3.NewPresignClient(service.Client)

	req, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(service.BucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expiration))

	if err != nil {
		slog.InfoContext(ctx, "Erro ao gerar URL pré-assinada: %v", err)
		return "", fmt.Errorf("Erro ao gerar URL pré-assinada: %w", err)
	}
	return req.URL, nil
}
