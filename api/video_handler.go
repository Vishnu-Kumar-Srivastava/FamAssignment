package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"ytvideofetcher/services"
)

func GetVideosHandler(c *gin.Context) {

	c.JSON(http.StatusOK, 1)
}

// FetchLatestVideos fetches latest videos from YouTube and stores in MongoDB
func Sync(c *gin.Context) {
	var ctx = c.Request.Context()
	fmt.Println("hello world")
	youtubeservice := services.NewYoutubeService()
	youtubeservice.GetVideos(ctx)

}
