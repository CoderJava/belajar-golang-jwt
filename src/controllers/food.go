package controllers

import (
	"belajar-golang-jwt/src/entities"
	"belajar-golang-jwt/src/helpers"
	"belajar-golang-jwt/src/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type FoodController struct {
	db    *gorm.DB
	error helpers.CustomError
}

func NewFoodController(db *gorm.DB, error helpers.CustomError) *FoodController {
	return &FoodController{db, error}
}

func (f FoodController) AddNewFoodItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createFoodItemBody := models.CreateFoodItemBody{}
		json.NewDecoder(r.Body).Decode(&createFoodItemBody)

		if len(createFoodItemBody.Name) < 3 {
			f.error.ApiError(w, http.StatusBadRequest, "Name should be at least 3 characters long!")
			return
		}

		if createFoodItemBody.Quantity == 0 {
			f.error.ApiError(w, http.StatusBadRequest, "Quantity shouldn't be zero!")
			return
		}

		if createFoodItemBody.SellingPrice == 0 {
			f.error.ApiError(w, http.StatusBadRequest, "Selling price shouldn't be zero!")
			return
		}

		food := entities.Food{
			Name:         createFoodItemBody.Name,
			Quantity:     createFoodItemBody.Quantity,
			SellingPrice: createFoodItemBody.SellingPrice,
		}
		if result := f.db.Create(&food); result.Error != nil {
			f.error.ApiError(w, http.StatusInternalServerError, "Failed to add new food item in database!")
			return
		}

		w.WriteHeader(http.StatusCreated)
		helpers.RespondWithJSON(
			w,
			models.SingleFoodItemResponse{
				ID:           food.ID,
				Name:         food.Name,
				Quantity:     food.Quantity,
				SellingPrice: food.SellingPrice,
			},
		)
	}
}

func (f FoodController) GetAllFoodItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		foodItems := []entities.Food{}

		if results := f.db.Find(&foodItems); results.Error != nil {
			f.error.ApiError(w, http.StatusInternalServerError, "Failed to fetch food items from database!")
			return
		}

		listFoodItem := []models.SingleFoodItemResponse{}
		for _, element := range foodItems {
			listFoodItem = append(listFoodItem, models.SingleFoodItemResponse{
				ID:           element.ID,
				Name:         element.Name,
				Quantity:     element.Quantity,
				SellingPrice: element.SellingPrice,
			})
		}

		allFoodItemResponse := models.AllFoodItemResponse{
			Data: listFoodItem,
		}
		helpers.RespondWithJSON(
			w,
			allFoodItemResponse,
		)
	}
}

func (f FoodController) GetSingleFoodItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathParams := mux.Vars(r)
		food := entities.Food{}

		result := f.db.
			Where("id = ?", pathParams["id"]).
			First(&food)
		if result.Error != nil || result.RowsAffected < 1 {
			f.error.ApiError(w, http.StatusNotFound, "Didn't find food item with id = "+pathParams["id"])
			return
		}

		helpers.RespondWithJSON(
			w,
			models.SingleFoodItemResponse{
				ID:           food.ID,
				Name:         food.Name,
				Quantity:     food.Quantity,
				SellingPrice: food.SellingPrice,
			},
		)
	}
}

func (f FoodController) DeleteFoodById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathParams := mux.Vars(r)
		food := entities.Food{}

		paramId := pathParams["id"]
		if result := f.db.Where("id = ?", paramId).First(&food); result.Error != nil || result.RowsAffected < 1 {
			f.error.ApiError(w, http.StatusNotFound, "Didn't find food item with id = "+paramId)
			return
		}

		if result := f.db.Delete(&food); result.Error != nil || result.RowsAffected < 1 {
			f.error.ApiError(w, http.StatusInternalServerError, "Failed to delete food from database!")
			return
		}

		helpers.RespondWithJSON(
			w,
			models.SingleFoodItemResponse{
				ID:           food.ID,
				Name:         food.Name,
				Quantity:     food.Quantity,
				SellingPrice: food.SellingPrice,
			},
		)
	}
}
