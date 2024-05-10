package model

type UserResponse struct {
	Id          int    `json:"id"`
	Username    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
}

type UserAddressResponse struct {
	Id         int    `json:"id"`
	DoorNo     string `json:"door_no"`
	StreetName string `json:"street_name"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	District   string `json:"district"`
}

type UserAddresRequest struct {
	DoorNo     string `json:"door_no"`
	StreetName string `json:"street_name"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	District   string `json:"district"`
	UserId     int    `json:"user_id"`
}

type UserRequest struct {
	Username    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
