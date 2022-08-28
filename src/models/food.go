package models

type AllFoodItemResponse struct {
	Data []SingleFoodItemResponse `json:"data"`
}

type SingleFoodItemResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Quantity     int8   `json:"quantity"`
	SellingPrice int    `json:"selling_price"`
}

type CreateFoodItemBody struct {
	Name         string `json:"name"`
	Quantity     int8   `json:"quantity"`
	SellingPrice int    `json:"selling_price"`
}
