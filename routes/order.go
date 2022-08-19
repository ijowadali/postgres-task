package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/controller"
)

func OrderRoute(router *gin.Engine) {
	router.GET("/orders", controller.GetOrder)
	router.POST("/orders", controller.CreateOrder)
}
