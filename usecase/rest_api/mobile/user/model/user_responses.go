package model

type UserResponse struct {
	Id          int    `json:"id"`
	Username    string `json:"user_name"`
	AddressId   int    `json:"address_id"`
	PhoneNumber string `json:"phone_number"`
}

type UserAddressRespons struct {
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
}

type UserRequest struct {
	Username    string            `json:"user_name"`
	PhoneNumber string            `json:"phone_number"`
	Address     *UserAddresRequest `json:"user_address"`
}
