package restapi

import (
	"fmt"
	"postgrest/usecase/rest_api/mobile/user/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(router fiber.Router) {

	//Register routes
	http.ReisterUserRoute(router)

	//use it for other services
	registerRouteMobileByVersion("v1", router, func(api fiber.Router) {})
}

type registerHandler func(api fiber.Router)

func registerRouteMobileByVersion(version string, router fiber.Router, handler registerHandler) {
	route := fmt.Sprintf("/%v/api/mobile", version)
	apiMobile := router.Group(route)
	handler(apiMobile)
}
