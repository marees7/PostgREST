package service

import (
	"context"
	"postgrest/pkg/cache"
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

	//Save user password in redis
	cache.RedisRestAPIClients.Client0.Set(context.Background(), user.Username, user.Password, 0)

	users, err := u.Repository.CreateUserRepo(user)
	if err != nil {
		return "", ErrorRepository
	}
	return users, NoError
}

func (u *UserService) AddUserAddressService(address model.UserAddresRequest) (string, pkgError.CustomError) {
	userAddress, err := u.Repository.AddUserAddresRepo(address)
	if err != nil {
		return userAddress, ErrorRepository
	}
	return userAddress, NoError
}

func (u *UserService) GetAllAddressesService() ([]model.UserAddressResponse, pkgError.CustomError) {
	users, err := u.Repository.GetAllAddressRepo()
	if err != nil {
		return nil, ErrorRepository
	}
	return users, NoError
}

func (u *UserService) GetAllUserAddressesService() ([]model.UserAddressResponse, pkgError.CustomError) {
	users, err := u.Repository.GetAllAddressRepo()
	if err != nil {
		return nil, ErrorRepository
	}
	return users, NoError
}
