package domain

import "postgrest/usecase/rest_api/mobile/user/model"

type UserRepositoryInterface interface {
	GetAllUsersRepo() ([]model.UserResponse, error)
}