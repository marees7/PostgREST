package http

import (
	helpers "postgrest/helpers/response"
	pkgError "postgrest/pkg/error"
	"postgrest/usecase/rest_api/mobile/user/model"
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

func (u *UserService) CreateUser(c *fiber.Ctx) error {

	//Get payload
	var payload = model.UserRequest{}
	err := c.BodyParser(&payload)
	if err != nil {
		return helpers.GenerateErrorResponseWithPKGError(c, pkgError.ErrorInvalidPayload.WithError(err), fiber.Map{"detail": err.Error()})
	}

	resp, customErr := u.Service.CreateUserService(payload)
	if !customErr.IsNoError() {
		helpers.GenerateCustomErrorResponseWithPKGError(c, customErr, resp)
	}

	return helpers.GenerateSuccessResponse(c, resp)
}

func (u *UserService) AddUserAddress(c *fiber.Ctx) error {

	//Get payload
	var payload = model.UserAddresRequest{}
	err := c.BodyParser(&payload)
	if err != nil {
		return helpers.GenerateErrorResponseWithPKGError(c, pkgError.ErrorInvalidPayload.WithError(err), fiber.Map{"detail": err.Error()})
	}

	resp, customErr := u.Service.AddUserAddressService(payload)
	if !customErr.IsNoError() {
		helpers.GenerateCustomErrorResponseWithPKGError(c, customErr, resp)
	}

	return helpers.GenerateCreatedSuccessResponse(c, resp)
}
