package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/controller"
)

func ItemRoute(router *gin.Engine) {
	router.GET("/item", controller.GetItem)
	router.POST("/item", controller.CreateItem)
}
