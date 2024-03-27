package daos

import (
	"context"
	
)

type IYoutubeDAO interface {
	GetVideos(ctx context.Context) (int, error)
}