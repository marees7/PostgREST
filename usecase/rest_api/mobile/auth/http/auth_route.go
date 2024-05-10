package http

import (
	"postgrest/usecase/rest_api/middleware"

	"github.com/gofiber/fiber/v2"
)

func ReisterAuthRoute(router fiber.Router) {
	version := middleware.NewVersion("auth", router)

	handler := NewAuthHandler()
	//define routes
	{
		version.RunMobile(func(apiWithVersion fiber.Router) {
			apiWithVersion.Post("login", handler.Login)
		}, "v1")
	}
}
