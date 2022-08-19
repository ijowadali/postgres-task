package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id         int    `json:"id" gorm:"primary_key"`
	OrderName  string `json:"order_name"`
	CustomerId int    `json:"customer_id"`
}
