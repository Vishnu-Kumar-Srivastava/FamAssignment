package services

import (
	"context"
	"ytvideofetcher/models"
)

type IYoutubeService interface {
	PullAndSaveVideos(ctx context.Context)  error
	GetVideos(ctx context.Context) ([]*models.Video, error)
}