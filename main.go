package main

import (
	"net/http"
	"ytvideofetcher/api"
	"ytvideofetcher/helpers"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8000"
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	//define routes here
	router.GET(helpers.Videos,api.Sync)
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
