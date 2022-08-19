package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/config"
	"github.com/jawad2020-web/postgress-task/models"
)

func GetOrder(ctx *gin.Context) {
	type Result struct {
		OrderName string `json:"order_name"`
		//Name      string `json:"name"`
	}
	list := []Result{}
	tx := config.DB.Table("orders").
		//Joins("INNER JOIN customers cc ON cc.id = orders.customer_id").
		Select("orders.order_name").
		Find(&list)
	//spg := paginate.New()
	// var result Result
	// var orders models.Order
	// //result := map[string]interface{}{}
	// model := config.DB.Model(&orders).Select("orders.order_name as OrderName, customers.name as Name").Joins("JOIN customers on orders.customer_id=customers.id").Find(&result)
	// sql := config.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Model(&orders).Select("orders.order_name, customers.name").Joins("JOIN customers on orders.customer_id=customers.id").Find(&result)
	// })
	println(list)
	// config.DB.Joins("JOIN customers on orders.customer_id=customers.id").Find(&orders)

	//var customer []models.Customers
	//model := config.DB.Joins("JOIN customers on orders.customer_id=customers.id").Find(&orders)
	// model := config.DB.Raw("SELECT name from customers").Scan(&result)
	ctx.JSON(200, tx)
}
func CreateOrder(ctx *gin.Context) {
	var order models.Order
	ctx.BindJSON(&order)
	config.DB.Create(&order)
	ctx.JSON(200, &order)
}
