package service

import (
	pkgError "postgrest/pkg/error"
	"postgrest/usecase/rest_api/mobile/user/domain"
	"postgrest/usecase/rest_api/mobile/user/model"
)

type UserService struct {
	Repository domain.UserRepositoryInterface
}

func NewUserService(repo domain.UserRepositoryInterface) *UserService {
	return &UserService{
		Repository: repo,
	}
}

func (u *UserService) GetAllUsers() ([]model.UserResponse, pkgError.CustomError) {
	users, err := u.Repository.GetAllUsersRepo()
	if err != nil {
		return nil, ErrorRepository
	}
	return users, NoError
}
