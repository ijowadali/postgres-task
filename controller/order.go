package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/config"
	"github.com/jawad2020-web/postgress-task/models"
	"github.com/morkid/paginate"
)

func GetOrder(ctx *gin.Context) {
	pg := paginate.New()
	var order []models.Order
	model := config.DB.Find(&order)
	ctx.JSON(200, pg.Response(model, ctx.Request, &order))
}
func CreateOrder(ctx *gin.Context) {
	var order models.Order
	ctx.BindJSON(&order)
	config.DB.Create(&order)
	ctx.JSON(200, &order)
}
