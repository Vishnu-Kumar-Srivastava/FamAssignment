package services

import (
	"context"
	"ytvideofetcher/models"
)

type IYoutubeService interface {
	PullAndSaveVideos(ctx context.Context)  error
	GetVideos(ctx context.Context, page, limit int) ([]*models.Video, error)
}