package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"ytvideofetcher/api"
	"ytvideofetcher/helpers"
	"ytvideofetcher/services"
	"context"
)

func main() {
	port := "8000"
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	// Define routes here
	router.GET(helpers.Videos, api.GetVideos)
	router.GET(helpers.Sync, api.Sync)

	// Start the server in a goroutine
	go func() {
		server := &http.Server{
			Addr:         ":" + port,
			Handler:      router,
			ReadTimeout:  360 * time.Second, // Set the read timeout
			WriteTimeout: 360 * time.Second, // Set the write timeout
		}

		// Start the server
		fmt.Println("Server is running on port", port)
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	// Start the cron job
	youtubeService := services.NewYoutubeService()
	gocron.Every(30).Seconds().Do(func() {
		// This function will be called every 30 seconds
		err := youtubeService.PullAndSaveVideos(context.Background())
		if err != nil {
			fmt.Println("Error pulling and saving videos:", err)
			// You might want to log the error to a file or external monitoring system
		}
	})
	<-gocron.Start()

	fmt.Println("Cron job started, but server is still running...")
}
