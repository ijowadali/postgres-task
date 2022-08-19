package models

import "gorm.io/gorm"

type Deliveries struct {
	gorm.Model
	Id                int `json:"id" gorm:"primary_key"`
	OrderItemId       int `json:"order_item_id"`
	DeliveredQuantity int `json:"delivered_quantity"`
}
