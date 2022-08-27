package router

import (
	"belajar-golang-jwt/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	userController := controllers.NewUserController(db)
	router.HandleFunc("/auth/signup", userController.SignupUser()).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", userController.LoginUser()).Methods(http.MethodPost)

	return router
}
