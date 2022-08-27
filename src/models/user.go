package models

type SignupUserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
