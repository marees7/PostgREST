package main

import (
	"log"
	"postgrest/config"
	"postgrest/pkg/cache"
	restapi "postgrest/usecase/rest_api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("postgREST api started...")

	// Get ENV
	errConfig := config.New("cmd/api/.env")
	if errConfig != nil {
		panic(errConfig)
	}

	//Redis connection
	err := cache.NewRedisRestAPIConnection()
	if err != nil {
		log.Fatal("Redis connection failed")
	}

	//Fiber setup
	app := fiber.New(fiber.Config{
		Prefork:   false,
		AppName:   "postgREST Demo",
		BodyLimit: (1 * 1024 * 1024), // Max 1MB (for image upload)
	})

	//Register route
	restapi.RegisterRoute(app)
	app.Listen(":8900")
}
