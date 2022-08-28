package controllers

import (
	"belajar-golang-jwt/src/entities"
	"belajar-golang-jwt/src/helpers"
	"belajar-golang-jwt/src/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type userController struct {
	db    *gorm.DB
	error helpers.CustomError
}

func NewUserController(db *gorm.DB, error helpers.CustomError) *userController {
	return &userController{db, error}
}

func (u userController) SignupUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := entities.User{}
		json.NewDecoder(r.Body).Decode(&user)

		if len(user.Name) < 3 {
			u.error.ApiError(w, http.StatusBadRequest, "Name should be at least 3 characters long!")
			return
		}

		if len(user.Username) < 3 {
			u.error.ApiError(w, http.StatusBadRequest, "Username should be at least 3 characters long!")
			return
		}

		if len(user.Email) < 3 {
			u.error.ApiError(w, http.StatusBadRequest, "Email should be at least 3 characters long!")
			return
		}

		if len(user.Password) < 3 {
			u.error.ApiError(w, http.StatusBadRequest, "Password should be at least 3 characters long!")
			return
		}

		if result := u.db.Create(&user); result.Error != nil {
			u.error.ApiError(w, http.StatusInternalServerError, "Failed to add new user n database!\n"+result.Error.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		helpers.RespondWithJSON(
			w,
			models.SignupUserResponse{
				Name:     user.Name,
				Email:    user.Email,
				Username: user.Username,
			},
		)
	}
}

func (u userController) LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := entities.User{}

		credentials := models.Credentials{}
		json.NewDecoder(r.Body).Decode(&credentials)

		if len(credentials.Id) < 3 {
			u.error.ApiError(w, http.StatusBadRequest, "Invalid username/email")
			return
		}

		if len(credentials.Password) < 3 {
			u.error.ApiError(w, http.StatusBadRequest, "Invalid password")
			return
		}

		result := u.db.
			Where("username = ? OR email = ?", credentials.Id, credentials.Id).
			First(&user)
		if result.Error != nil || result.RowsAffected < 1 {
			u.error.ApiError(w, http.StatusNotFound, "Invalid username/email, please signup!")
			return
		}

		if user.Password != credentials.Password {
			u.error.ApiError(w, http.StatusNotFound, "Invalid credentials!")
			return
		}

		payload := helpers.Payload{
			Username: user.Username,
			Email:    user.Email,
			Id:       user.ID,
		}

		token, err := helpers.GenerateJwtToken(payload)
		if err != nil {
			u.error.ApiError(w, http.StatusInternalServerError, "Failed to generate JWT token!")
			return
		}

		helpers.RespondWithJSON(
			w,
			models.LoginUserResponse{
				Token: token,
			},
		)
	}
}
