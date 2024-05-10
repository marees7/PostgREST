package model

type LoginRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId      int    `json:"user_id"`
	Username    string `json:"user_name"`
	Token       string `json:"token"`
	ExpiryTime  int    `json:"expiry_time"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
}
