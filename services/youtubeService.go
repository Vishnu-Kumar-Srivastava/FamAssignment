package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ytvideofetcher/daos"
	"ytvideofetcher/models"
)

type YoutubeService struct {
}

// Delete(ctx context.Context, accountId string) error
// UpsertMany(ctx context.Context, orgId string, accounts []*models.Account) error
func NewYoutubeService() IYoutubeService {
	return &YoutubeService{}
}

func (s *YoutubeService) GetVideos(ctx context.Context) {
	apiKey := "AIzaSyDwUN9D85UP-Y0cNurozBoONoCP2Vj2eHg"
	query := "football"

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?key=%s&q=%s&type=video&part=snippet", apiKey, query)

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	videos := &models.Response{}

	err = json.Unmarshal([]byte(string(body)), &videos)
	if err != nil {

		fmt.Println("Error found!")
	}
	dao := daos.NewYtVideoDAO()
	dao.UpsertVideos(ctx, videos)
	

	x,err := dao.GetVideos(ctx)
	fmt.Println(x)

	// Print the response body (or do something else with it)
	// fmt.Println(videos)
}
