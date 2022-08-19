package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/config"
	"github.com/jawad2020-web/postgress-task/models"
)

func GetDeliveries(ctx *gin.Context) {
	var deliveries []models.Deliveries
	config.DB.Find(&deliveries)
	ctx.JSON(200, &deliveries)
}
func CreateDeliveries(ctx *gin.Context) {
	var deliveries models.Deliveries
	ctx.BindJSON(&deliveries)
	config.DB.Create(&deliveries)
	ctx.JSON(200, &deliveries)
}
