package main

import (
	"net/http"
	"ytvideofetcher/api"
	"ytvideofetcher/helpers"
	"github.com/jasonlvhit/gocron"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ytvideofetcher/services"
	"fmt"
	"context"
)

func main() {
	port := "8000"
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	//define routes here
	router.GET(helpers.Videos,api.GetVideos)
	router.GET(helpers.Sync,api.Sync)

	youtubeService := services.NewYoutubeService()
	gocron.Every(30).Seconds().Do(func() {
		// This function will be called every 30 seconds
		err := youtubeService.PullAndSaveVideos(context.Background())
		if err != nil {
			fmt.Println("Error pulling and saving videos:", err)
		}
	})
	<-gocron.Start()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  360 * time.Second, // Set the read timeout
		WriteTimeout: 360 * time.Second, // Set the write timeout
	}

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

	


	
}

