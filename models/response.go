package models

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type CheckTokenResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
