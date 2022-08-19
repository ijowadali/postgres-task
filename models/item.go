package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Id           int     `json:"id" gorm:"primary_key"`
	OrderId      int     `json:"order_id"`
	PricePerUnit float32 `json:"price_per_unit"`
	Quantity     int     `json:"quantity"`
	Product      string  `json:"product"`
}
