package repository

import (
	"encoding/json"
	"net/http"
	"postgrest/config"
	"postgrest/usecase/rest_api/mobile/auth/domain"
	"postgrest/usecase/rest_api/mobile/user/model"
)

func NewUserRepository() domain.AuthRepositoryInterface {
	return &AuthRepoImpl{}
}

type AuthRepoImpl struct {
}

func (a *AuthRepoImpl) GetUserByUsername(name string) (*model.UserResponse, error) {
	var response []model.UserResponse
	var userResp = &model.UserResponse{}
	host := config.Configuration.Get("POSTGREST_HOST")
	port := config.Configuration.Get("POSTGREST_PORT")
	apiURI := host + ":" + port + "/get_all_users" + "?user_name=eq." + name
	http.Header.Add(http.Header{}, "Authorization", "Bearer "+config.Configuration.Get("POSTGREST_TOKEN"))
	resp, err := http.Get(apiURI)
	if err != nil {
		return userResp, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return userResp, err
	}
	for _, v := range response {
		userResp.Email = v.Email
		userResp.Gender = v.Gender
		userResp.Id = v.Id
		userResp.Username = v.Username
		userResp.PhoneNumber = v.PhoneNumber
	}
	return userResp, nil
}
