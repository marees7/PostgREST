package service

import (
	"context"
	"postgrest/pkg/cache"
	pkgError "postgrest/pkg/error"
	"postgrest/pkg/jwt"
	"postgrest/usecase/rest_api/mobile/auth/domain"
	"postgrest/usecase/rest_api/mobile/auth/model"
)

type AuthService struct {
	Repository domain.AuthRepositoryInterface
}

func NewAuthService(repo domain.AuthRepositoryInterface) *AuthService {
	return &AuthService{
		Repository: repo,
	}
}

func (a *AuthService) LoginService(req model.LoginRequest) (*model.LoginResponse, pkgError.CustomError) {
	var response *model.LoginResponse
	savedPassword := cache.RedisRestAPIClients.Client0.Get(context.Background(), req.Username).Val()
	if savedPassword != req.Password {
		return response, ErrorInvalidPassword
	}
	resp, err := a.Repository.GetUserByUsername(req.Username)
	if err != nil {
		return response, ErrorRepository
	}
	token, err := jwt.GenerateJWTToken(req.Username)
	if err != nil {
		return response, ErrorRepository
	}
	response = &model.LoginResponse{
		UserId:      resp.Id,
		Username:    resp.Username,
		Email:       resp.Email,
		Gender:      resp.Gender,
		PhoneNumber: resp.PhoneNumber,
		Token:       token,
		ExpiryTime:  1 * 3600,
	}
	return response, NoError
}
