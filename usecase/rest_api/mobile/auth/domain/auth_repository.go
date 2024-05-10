package domain

import "postgrest/usecase/rest_api/mobile/user/model"

type AuthRepositoryInterface interface {
	GetUserByUsername(name string) (*model.UserResponse, error)
}
