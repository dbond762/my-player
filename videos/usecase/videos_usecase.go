package usecase

import (
	models "github.com/dbond762/my-player/videos"
)

type VideosUsecase interface {
	Search(query string) ([]models.Video, error)
}

type videosUsecase struct {

}

func NewVideosUsecase() VideosUsecase {
	return &videosUsecase{}
}

func (v *videosUsecase) Search(query string) ([]models.Video, error) {

}
