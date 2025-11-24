package config

import (
	"github.com/Avinashbhat1199/Golang-Basics/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhosty:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic("db error ")
	}
	DB = db
	DB.AutoMigrate(&models.User{})
}
