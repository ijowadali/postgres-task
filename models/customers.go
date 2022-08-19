package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Id        int    `json:"id" gorm:"primary_key"`
	UserId    string `json:"user_id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	CompanyId int    `json:"company_id"`
}
