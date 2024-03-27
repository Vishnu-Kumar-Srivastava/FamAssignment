package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"
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

func (s *YoutubeService) PullAndSaveVideos(ctx context.Context) error {
	apiKey := "AIzaSyDwUN9D85UP-Y0cNurozBoONoCP2Vj2eHg"
	query := "football"
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).Format("2006-01-02T15:04:05Z")

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?key=%s&q=%s&type=video&part=snippet&order=date&publishedAfter=%s", apiKey, query, sevenDaysAgo)

	// Create a new GET request

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	videos := &models.Response{}

	err = json.Unmarshal([]byte(string(body)), &videos)
	if err != nil {

		fmt.Println("Error found!")
	}
	dao := daos.NewYtVideoDAO()
	dao.UpsertVideos(ctx, videos)
	count := 0
	for {
		if count >= 6 {
			break
		}
		count++
		pageToken := videos.NextPageToken
		if pageToken == "" {
			break
		}
		url = fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?key=%s&q=%s&type=video&part=snippet&order=date&publishedAfter=%s&pageToken=%s", apiKey, query, sevenDaysAgo, pageToken)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}

		// Send the request
		client = &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return err
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return err
		}

		err = json.Unmarshal([]byte(string(body)), &videos)
		if err != nil {

			fmt.Println("Error found!")
		}
		dao := daos.NewYtVideoDAO()
		dao.UpsertVideos(ctx, videos)

	}

	return nil

	// Print the response body (or do something else with it)
	// fmt.Println(videos)
}

func (s *YoutubeService) GetVideos(ctx context.Context) ([]*models.Video, error) {
	dao := daos.NewYtVideoDAO()
	videos, err := dao.GetVideos(ctx)
	if err != nil {
		return nil, err
	}

	// Sort videos based on PublishedAt
	sort.Slice(videos, func(i, j int) bool {
		return videos[i].Snippet.PublishedAt.After(videos[j].Snippet.PublishedAt)
	})

	return videos, nil

}
