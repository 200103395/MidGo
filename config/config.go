package config

import (
	"MiddleProject/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "almas:almas@12@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = d
	d.AutoMigrate(&models.User{})
}

func getDB() *gorm.DB {
	return DB
}
