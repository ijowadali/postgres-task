package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Id          int    `json:"id" gorm:"primary_key"`
	CompanyId   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
}
