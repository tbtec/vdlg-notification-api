package repository

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/infra/database"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type IVideoRepository interface {
	Create(ctx context.Context, order *model.Video) error
	Find(ctx context.Context) ([]model.Video, error)
	FindOne(ctx context.Context, id string) (*model.Video, error)
	Update(ctx context.Context, order *model.Video) error
}

type VideoRepository struct {
	database database.RDBMS
}

func NewVideoRepository(database database.RDBMS) IVideoRepository {
	return &VideoRepository{
		database: database,
	}
}

func (repository *VideoRepository) Create(ctx context.Context, order *model.Video) error {

	result := repository.database.DB.WithContext(ctx).Create(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *VideoRepository) Find(ctx context.Context) ([]model.Video, error) {
	videos := []model.Video{}

	result := repository.database.DB.WithContext(ctx).Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}

	return videos, nil
}

func (repository *VideoRepository) FindOne(ctx context.Context, id string) (*model.Video, error) {
	video := &model.Video{}

	result := repository.database.DB.WithContext(ctx).Where("video_id = ?", id).First(&video)

	if result.Error != nil {
		return nil, result.Error
	}

	return video, nil
}

func (repository *VideoRepository) Update(ctx context.Context, video *model.Video) error {

	slog.InfoContext(ctx, "Updating video", slog.Any("video", video))

	result := repository.database.DB.WithContext(ctx).Save(&video)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
