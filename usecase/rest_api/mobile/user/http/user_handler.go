package http

import (
	helpers "postgrest/helpers/response"
	"postgrest/usecase/rest_api/mobile/user/repository"
	"postgrest/usecase/rest_api/mobile/user/service"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	Service *service.UserService
}

func NewUserHandler() *UserService {
	return &UserService{
		Service: service.NewUserService(
			repository.NewUserRepository(),
		),
	}
}

func (u *UserService) GetUsers(c *fiber.Ctx) error {

	resp, err := u.Service.GetAllUsers()
	if !err.IsNoError() {
		helpers.GenerateCustomErrorResponseWithPKGError(c, err, resp)
	}

	return helpers.GenerateSuccessResponse(c, resp)
}
