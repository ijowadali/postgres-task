package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/controller"
)

func CompanyRoute(router *gin.Engine) {
	router.GET("/company", controller.GetCompany)
	router.POST("/company", controller.CreateCompany)
}
