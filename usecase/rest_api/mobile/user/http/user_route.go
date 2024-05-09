package http

import (
	"postgrest/usecase/rest_api/middleware"

	"github.com/gofiber/fiber/v2"
)

func ReisterUserRoute(router fiber.Router) {
	version := middleware.NewVersion("users", router)

	handler := NewUserHandler()
	//define routes
	{
		version.RunMobile(func(apiWithVersion fiber.Router) {
			apiWithVersion.Post("add-address", handler.AddUserAddress)
			apiWithVersion.Get("", handler.GetUsers)
		}, "v1")
	}
}
