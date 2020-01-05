package main

import (
	"fmt"
	"log"
	"os"

	userModels "github.com/dbond762/my-player/user"
	userRepo "github.com/dbond762/my-player/user/repository"
	videosModels "github.com/dbond762/my-player/videos"
	videosHttpDeliver "github.com/dbond762/my-player/videos/delivery/http"
	videosRepo "github.com/dbond762/my-player/videos/repository"
	videosUcase "github.com/dbond762/my-player/videos/usecase"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	app.Favicon("./frontend/build/favicon.ico")

	app.HandleDir("/", "./frontend/build", iris.DirOptions{
		IndexName: "/index.html",
		Gzip:      false,
		ShowList:  false,
	})

	db, err := gorm.Open("mysql", "root:123456@/my_player?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
	}()

	db.AutoMigrate(&videosModels.Video{}, &videosModels.Query{}, &userModels.User{})

	videosRepo.NewDBVideosRepository(db)
	redisVideosRepo := videosRepo.NewRedisVideosRepository("redis://youtube_player:@localhost:6379/0")
	httpVideosRepo := videosRepo.NewHttpVideosRepository()

	userRepo.NewDBUserRepository(db)

	videosUsecase := videosUcase.NewVideosUsecase(redisVideosRepo, httpVideosRepo)

	api := app.Party("/api")
	videosHttpDeliver.NewVideosHandler(api, videosUsecase)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := app.Run(iris.Addr(addr), iris.WithoutServerError(iris.ErrServerClosed)); err != nil {
		log.Fatal(err)
	}
}
