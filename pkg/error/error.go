package pkgError

import "github.com/gofiber/fiber/v2"

type CustomError struct {
	Message  string
	Code     string
	HttpCode int
	Err      error
	Detail   interface{}
}

func (e CustomError) Error() string {
	if e.Err == nil {
		return e.Message
	}
	return e.Err.Error()
}

func (e CustomError) IsNoError() bool {
	return e == CustomError{}
}

var (
	ErrorInvalidToken CustomError = CustomError{Code: "authentication_failed", Message: "Invalid Token", HttpCode: fiber.StatusUnauthorized}
)
