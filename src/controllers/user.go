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
	db *gorm.DB
}

var error = helpers.CustomError{}

func NewUserController(db *gorm.DB) *userController {
	return &userController{db}
}

func (u userController) SignupUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := entities.User{}
		json.NewDecoder(r.Body).Decode(&user)

		if len(user.Name) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Name should be at least 3 characters long!")
			return
		}

		if len(user.Username) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Username should be at least 3 characters long!")
			return
		}

		if len(user.Email) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Email should be at least 3 characters long!")
			return
		}

		if len(user.Password) < 3 {
			error.ApiError(w, http.StatusBadRequest, "Password should be at least 3 characters long!")
			return
		}

		if result := u.db.Create(&user); result.Error != nil {
			error.ApiError(w, http.StatusInternalServerError, "Failed to add new user n database!\n"+result.Error.Error())
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
