package usecase

import (
	models "github.com/dbond762/my-player/videos"
	"github.com/dbond762/my-player/videos/repository"
)

type VideosUsecase interface {
	Search(query string) ([]models.Video, error)
}

type videosUsecase struct {
	redisVideosRepo *repository.RedisVideosRepository
	httpVideosRepo *repository.HttpVideosRepository
}

func NewVideosUsecase(rv *repository.RedisVideosRepository, hv *repository.HttpVideosRepository) VideosUsecase {
	return &videosUsecase{
		redisVideosRepo: rv,
		httpVideosRepo: hv,
	}
}

func (v *videosUsecase) Search(query string) ([]models.Video, error) {
	videos, err := v.redisVideosRepo.GetVideosFromCache(query)
	if err == repository.VideosNotInCache {
		videos, err = v.httpVideosRepo.SearchVideosByYoutubeAPI(query)
		if err != nil {
			return []models.Video{}, err
		}

		if err := v.redisVideosRepo.AddVideosToCache(query, videos); err != nil {
			return []models.Video{}, err
		}
	} else if err != nil {
		return []models.Video{}, err
	}

	return videos, nil
}
