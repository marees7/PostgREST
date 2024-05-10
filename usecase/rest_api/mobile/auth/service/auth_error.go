package service

import (
	pkgError "postgrest/pkg/error"

	"github.com/gofiber/fiber"
)

type ErrorDetail struct {
	Title     string      `json:"title"`
	Content   string      `json:"content"`
	Detail    string      `json:"detail"`
	EventName interface{} `json:"event_name,omitempty"`
}

var (
	NoError              pkgError.CustomError = pkgError.CustomError{}
	ErrorRepository      pkgError.CustomError = pkgError.CustomError{Code: "general_error", Message: "general error", HttpCode: fiber.StatusInternalServerError}
	ErrorInvalidPassword pkgError.CustomError = pkgError.CustomError{Code: "invalid_password", Message: "password is invalid", HttpCode: fiber.StatusBadRequest}
)
