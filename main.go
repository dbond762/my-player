package main

import (
	"fmt"
	"log"
	"os"

	videosHttpDeliver "github.com/dbond762/my-player/videos/delivery/http"
	videosRepo "github.com/dbond762/my-player/videos/repository"
	videosUcase "github.com/dbond762/my-player/videos/usecase"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func init() {
	env := os.Getenv("env")
	envFile := fmt.Sprintf("%s.env", env)
	if err := godotenv.Load(envFile); err != nil {
		log.Print("Environment file not found")
	}
}

func main() {
	app := iris.Default()

	redisVideosRepo := videosRepo.NewRedisVideosRepository("redis://youtube_player:@localhost:6379/0")
	httpVideosRepo := videosRepo.NewHttpVideosRepository()

	videosUsecase := videosUcase.NewVideosUsecase(redisVideosRepo, httpVideosRepo)

	videosHttpDeliver.NewVideosHandler(app, videosUsecase)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := app.Run(iris.Addr(addr), iris.WithoutServerError(iris.ErrServerClosed)); err != nil {
		log.Fatal(err)
	}
}
