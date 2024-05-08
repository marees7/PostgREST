package repository

import (
	"encoding/json"
	"net/http"
	"postgrest/config"
	"postgrest/usecase/rest_api/mobile/user/domain"
	"postgrest/usecase/rest_api/mobile/user/model"
)

func NewUserRepository() domain.UserRepositoryInterface {
	return &UserRepoImpl{}
}

type UserRepoImpl struct {
}

func (u *UserRepoImpl) GetAllUsersRepo() ([]model.UserResponse, error) {
	users := []model.UserResponse{}

	host := config.Configuration.Get("POSTGREST_HOST")
	port := config.Configuration.Get("POSTGREST_PORT")
	apiURI := host + ":" + port + "/get_all_users"
	http.Header.Add(http.Header{}, "Authorization", "Bearer "+config.Configuration.Get("POSTGREST_TOKEN"))
	resp, err := http.Get(apiURI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
