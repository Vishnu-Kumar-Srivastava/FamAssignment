package services
import(
	"context"
)

type IYoutubeService interface {
	GetVideos(ctx context.Context)
}