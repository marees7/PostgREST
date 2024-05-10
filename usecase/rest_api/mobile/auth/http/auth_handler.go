package http

import (
	helpers "postgrest/helpers/response"
	pkgError "postgrest/pkg/error"
	"postgrest/usecase/rest_api/mobile/auth/model"
	"postgrest/usecase/rest_api/mobile/auth/repository"
	"postgrest/usecase/rest_api/mobile/auth/service"

	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	Service *service.AuthService
}

func NewAuthHandler() *AuthService {
	return &AuthService{
		Service: service.NewAuthService(
			repository.NewUserRepository(),
		),
	}
}

func (u *AuthService) Login(c *fiber.Ctx) error {

	//Get payload
	var payload = model.LoginRequest{}
	err := c.BodyParser(&payload)
	if err != nil {
		return helpers.GenerateErrorResponseWithPKGError(c, pkgError.ErrorInvalidPayload.WithError(err), fiber.Map{"detail": err.Error()})
	}

	resp, customErr := u.Service.LoginService(payload)
	if !customErr.IsNoError() {
		return helpers.GenerateCustomErrorResponseWithPKGError(c, customErr, resp)
	}

	return helpers.GenerateSuccessResponse(c, resp)
}
