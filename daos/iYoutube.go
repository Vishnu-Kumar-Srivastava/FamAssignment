package daos

import (
	"context"
	"ytvideofetcher/models"
)

type IYtVideoDAO interface {
	UpsertVideos(context.Context, *models.Response) error
	GetVideos(ctx context.Context) ([]*models.Video, error)
}
