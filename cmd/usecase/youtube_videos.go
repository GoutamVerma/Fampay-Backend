package usecase

import (
	"log"
	"net/http"
	"time"

	config "github.com/GoutamVerma/FamPay-Backend/config"
	models "github.com/GoutamVerma/FamPay-Backend/utils"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func StartYouTubeAPICall(query string, maxResults int64) {
	apiKeys := config.ReadYouTubeAPIKeys()
	apiKeyIndex := 0
	interval := config.FetchInternval()
	go func() {
		ticker := time.NewTicker(time.Duration(interval) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				go func() {
					apiKey := apiKeys[apiKeyIndex]
					err := fetchAndStoreVideos(apiKey, query, maxResults)
					if err != nil {
						log.Println("Error fetching and storing videos:", err)
					}

					// Rotate to the next API key
					apiKeyIndex = (apiKeyIndex + 1) % len(apiKeys)
				}()
			}
		}
	}()
}

func fetchAndStoreVideos(apiKey string, query string, maxResults int64) error {
	youtubeService, err := youtube.New(&http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	})
	if err != nil {
		return err
	}

	videos, err := fetchLatestVideos(youtubeService, query, maxResults)
	if err != nil {
		return err
	}

	// Store the fetched videos in the database
	for _, video := range videos {
		go func(v models.Video) {
			err := sql.AddVideos(v)
			if err != nil {
				log.Println("Error storing video:", err)
			}
		}(video)
	}

	return nil
}

/*
fetchLatestVideos fetches the latest videos from YouTube based on the given query and maximum number of results.
It uses the provided YouTube service to make the API call and retrieve the videos./
The function returns a slice of models.Video and an error if any occurred.
*/
func fetchLatestVideos(service *youtube.Service, query string, maxResults int64) ([]models.Video, error) {
	// Create a search call with the specified parameters
	call := service.Search.List([]string{"id", "snippet"}).
		Q(query).
		Type("video").
		Order("date").
		MaxResults(maxResults)

	// Execute the API call
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var videos []models.Video
	for _, item := range response.Items {
		// Parse the publishedAt field into a time.Time object
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			return nil, err
		}

		// Create a models.Video object using the retrieved data
		video := models.Video{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			PublishedAt: publishedAt,
			Thumbnails:  item.Snippet.Thumbnails.Default.Url,
		}
		videos = append(videos, video)
	}

	return videos, nil
}
