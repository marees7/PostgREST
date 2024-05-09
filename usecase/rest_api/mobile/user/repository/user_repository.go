package repository

import (
	"bytes"
	"encoding/json"
	"errors"
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

func (u *UserRepoImpl) AddUserAddresRepo(address model.UserAddresRequest) (string, error) {

	host := config.Configuration.Get("POSTGREST_HOST")
	port := config.Configuration.Get("POSTGREST_PORT")
	apiURI := host + ":" + port + "/user_address"
	userAddressReq, err := json.Marshal(address)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(apiURI, "application/json", bytes.NewBuffer(userAddressReq))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return "", errors.New("invalid status code")
	}
	return UserCreated, nil
}
