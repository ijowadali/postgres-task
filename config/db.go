package config

import (
	"github.com/jawad2020-web/postgress-task/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Deliveries{})
	db.AutoMigrate(&models.Customers{})
	db.AutoMigrate(&models.Company{})
	DB = db
}
