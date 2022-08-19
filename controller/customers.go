package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/config"
	"github.com/jawad2020-web/postgress-task/models"
)

func GetCustomers(ctx *gin.Context) {
	var customers []models.Customers
	config.DB.Find(&customers)
	ctx.JSON(200, &customers)
}
func CreateCustomers(ctx *gin.Context) {
	var customers models.Customers
	ctx.BindJSON(&customers)
	config.DB.Create(&customers)
	ctx.JSON(200, &customers)
}
