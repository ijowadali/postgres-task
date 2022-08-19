package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/config"
	"github.com/jawad2020-web/postgress-task/models"
)

func GetCompany(ctx *gin.Context) {
	var company []models.Company
	config.DB.Find(&company)
	ctx.JSON(200, &company)
}
func CreateCompany(ctx *gin.Context) {
	var company models.Company
	ctx.BindJSON(&company)
	config.DB.Create(&company)
	ctx.JSON(200, &company)
}
