package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/snow-dev/super_blog/api/controllers"
	"github.com/snow-dev/super_blog/api/seed"
	"log"
	"os"
)

var server = controllers.Server{}

func init() {
	// load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("SAD .env file not found")
	}
}

func Run() {
	var err error
	var address string

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
		address = ":" + os.Getenv("API_PORT")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	seed.Load(server.DB)
	server.Run(address)
}
