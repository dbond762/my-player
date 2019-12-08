package http

import (
	videosUcase "github.com/dbond762/my-player/videos/usecase"
	"github.com/kataras/iris/v12"
	models "github.com/dbond762/my-player/videos"
)

type HttpVideosHandler struct {
	VUsecase videosUcase.VideosUsecase
}

func (v *HttpVideosHandler) Search(ctx iris.Context) {
	query := ctx.Params().GetString("query")

	videos, err := v.VUsecase.Search(query)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	i, err := ctx.JSON(videos)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
}
