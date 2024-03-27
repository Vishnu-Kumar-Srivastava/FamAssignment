package daos

import (
	"context"
	"ytvideofetcher/models"
)

type IYtVideoDAO interface {
	UpsertVideos(context.Context, *models.Response) error
	GetVideos(context.Context) ([]*models.Video,error)
}
