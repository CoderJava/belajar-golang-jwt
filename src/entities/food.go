package entities

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name         string
	Quantity     int8
	SellingPrice int `gorm:"column:selling_price"`
}
