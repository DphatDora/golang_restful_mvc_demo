package config

import (
	"fmt"
	"go_restful_mvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:ducphat1708@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	fmt.Println("Connected to the database successfully")
	DB = db
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
}