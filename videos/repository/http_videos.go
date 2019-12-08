package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	models "github.com/dbond762/my-player/videos"
)

const (
	ApiKey     = "AIzaSyBVJgyC-x6CsM-hPCYY10VfOnGOKksDK8U"
	maxResults = 25
)

type HttpVideosRepository interface {
	SearchVideosByYoutubeAPI(query string) ([]models.Video, error)
}

type httpVideosRepository struct {
}

func NewHttpVideosRepository() *httpVideosRepository {
	return &httpVideosRepository{}
}

type ApiSearchResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				High struct {
					URL string `json:"url"`
				} `json:"high"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

func (h *httpVideosRepository) SearchVideosByYoutubeAPI(query string) ([]models.Video, error) {
	baseUrl := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=%d&q=%s&key=%s", maxResults, query, ApiKey)
	resp, err := http.Get(baseUrl)
	if err != nil {
		log.Print(err)
		return []models.Video{}, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Print(err)
		}
	}()

	decoder := json.NewDecoder(resp.Body)
	apiResp := new(ApiSearchResponse)
	if err := decoder.Decode(apiResp); err != nil {
		log.Print(err)
		return []models.Video{}, err
	}

	videos := make([]models.Video, len(apiResp.Items))
	for i, video := range apiResp.Items {
		v := models.Video{
			ID:          video.ID.VideoID,
			Title:       video.Snippet.Title,
			PubDate:     video.Snippet.PublishedAt,
			Description: video.Snippet.Description,
			Thumbnail:   video.Snippet.Thumbnails.High.URL,
		}
		videos[i] = v
	}

	return videos, nil
}
