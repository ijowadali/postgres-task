package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/controller"
)

func CustomersRoute(router *gin.Engine) {
	router.GET("/customers", controller.GetCustomers)
	router.POST("/customers", controller.CreateCustomers)
}
