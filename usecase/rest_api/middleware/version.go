package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type version struct {
	service string
	router  fiber.Router
}

func NewVersion(service string, router fiber.Router) version {
	return version{
		service: service,
		router:  router,
	}
}

func (v version) RunMobile(handler func(apiWithVersion fiber.Router), versions ...string) {
	for _, version := range versions {
		groupString := fmt.Sprintf("/%v/api/mobile/%v", version, v.service)
		apiGroup := v.router.Group(groupString)

		handler(apiGroup)
	}
}
