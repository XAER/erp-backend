package models

type RegisterRequestBody struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginRequestBody struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
