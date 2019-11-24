package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("env")
	envFile := fmt.Sprintf("%s.env", env)
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Environment file not found")
	}
}

func main() {

}
