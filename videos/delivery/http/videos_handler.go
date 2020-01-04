package http

import (
	videosUcase "github.com/dbond762/my-player/videos/usecase"
	"github.com/kataras/iris/v12"
)

type VideosHandler struct {
	VUsecase videosUcase.VideosUsecase
}

func NewVideosHandler(app iris.Party, vu videosUcase.VideosUsecase) {
	handler := &VideosHandler{
		VUsecase: vu,
	}

	app.Get("/search/", handler.Search)
	app.Get("/search/{query}", handler.Search)
}

func (v *VideosHandler) Search(ctx iris.Context) {
	query := ctx.Params().Get("query")

	videos, err := v.VUsecase.Search(query)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	if _, err := ctx.JSON(videos); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
}
