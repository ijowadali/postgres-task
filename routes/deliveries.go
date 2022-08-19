package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/controller"
)

func DeliveriesRoute(router *gin.Engine) {
	router.GET("/deliveries", controller.GetDeliveries)
	router.POST("/deliveries", controller.CreateDeliveries)
}
