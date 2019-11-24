package main

import (
	"fmt"
	"log"
	"os"

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

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := app.Run(iris.Addr(addr), iris.WithoutServerError(iris.ErrServerClosed)); err != nil {
		log.Fatal(err)
	}
}
