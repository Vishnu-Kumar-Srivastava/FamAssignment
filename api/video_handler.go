package api

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"ytvideofetcher/services"
)


// FetchLatestVideos fetches latest videos from YouTube and stores in MongoDB
func Sync(c *gin.Context) {
	var ctx = c.Request.Context()
	fmt.Println("hello world")
	youtubeservice := services.NewYoutubeService()
	err := youtubeservice.PullAndSaveVideos(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync videos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Videos synced successfully"})
}



func GetVideos(c *gin.Context) {
	var ctx = c.Request.Context()
	youtubeservice := services.NewYoutubeService()

	// Extract page and limit from query parameters (default to 1 and 10 if not provided)
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit <= 0 {
		limit = 10
	}

	videos, err := youtubeservice.GetVideos(ctx, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get videos"})
		return
	}
	c.JSON(http.StatusOK, videos)
}
