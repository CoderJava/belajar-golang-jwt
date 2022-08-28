package router

import (
	"belajar-golang-jwt/src/controllers"
	"belajar-golang-jwt/src/helpers"
	"belajar-golang-jwt/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	error := helpers.CustomError{}

	userController := controllers.NewUserController(db, error)
	router.HandleFunc("/auth/signup", userController.SignupUser()).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", userController.LoginUser()).Methods(http.MethodPost)

	foodController := controllers.NewFoodController(db, error)
	router.HandleFunc("/food", middlewares.CheckAuth(foodController.AddNewFoodItem())).Methods(http.MethodPost)
	router.HandleFunc("/food", middlewares.CheckAuth(foodController.GetAllFoodItems())).Methods(http.MethodGet)
	router.HandleFunc("/food/{id}", middlewares.CheckAuth(foodController.GetSingleFoodItem())).Methods(http.MethodGet)
	router.HandleFunc("/food/{id}", middlewares.CheckAuth(foodController.DeleteFoodById())).Methods(http.MethodDelete)

	return router
}
