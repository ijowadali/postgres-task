package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jawad2020-web/postgress-task/config"
	"github.com/jawad2020-web/postgress-task/routes"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	config.Connect()
	routes.OrderRoute(router)
	routes.ItemRoute(router)
	routes.DeliveriesRoute(router)
	routes.CustomersRoute(router)
	routes.CompanyRoute(router)
	router.Run(":8082")
}
