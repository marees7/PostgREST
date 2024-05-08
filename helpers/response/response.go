package helpers

import (
	pkgError "postgrest/pkg/error"

	"github.com/gofiber/fiber/v2"
)

func GenerateCustomErrorResponseWithPKGError(c *fiber.Ctx, customError pkgError.CustomError, data interface{}) error {
	return c.Status(customError.HttpCode).JSON(ErrorResponse{
		Success: false,
		Message: customError.Message,
		Code:    customError.Code,
		Data:    data,
		Error:   customError.Detail,
	})
}

func GenerateSuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Message: "success",
		Data:    data,
	})

}
