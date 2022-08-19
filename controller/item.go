package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/config"
	"github.com/jawad2020-web/postgress-task/models"
)

func GetItem(ctx *gin.Context) {
	var item []models.Item
	config.DB.Find(&item)
	ctx.JSON(200, &item)
}
func CreateItem(ctx *gin.Context) {
	var item models.Item
	ctx.BindJSON(&item)
	config.DB.Create(&item)
	ctx.JSON(200, &item)
}
