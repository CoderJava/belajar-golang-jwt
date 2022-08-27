package models

type SignupUserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Credentials struct {
	Id       string
	Password string
}

type LoginUserResponse struct {
	Token string `json:"token"`
}
