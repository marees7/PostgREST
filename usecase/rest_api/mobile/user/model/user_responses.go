package model

type UserResponse struct {
	Id          int    `json:"id"`
	Username    string `json:"user_name"`
	AddressId   int    `json:"address_id"`
	PhoneNumber string `json:"phone_number"`
}
