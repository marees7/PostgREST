package service

import (
	"fmt"
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

func (u *UserService) CreateUserService(user model.UserRequest) (string, pkgError.CustomError) {
	users, err := u.Repository.CreateUserRepo(user)
	if err != nil {
		return "", ErrorRepository
	}
	return users, NoError
}

func (u *UserService) AddUserAddressService(address model.UserAddresRequest) (string, pkgError.CustomError) {
	userAddress, err := u.Repository.AddUserAddresRepo(address)
	fmt.Println("user add ::", userAddress, err)
	if err != nil {
		return userAddress, ErrorRepository
	}
	return userAddress, NoError
}
