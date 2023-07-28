package config

import (
	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := "root"
	password := "Kanarkaan73"
	host := "tcp(127.0.0.1:3306)"
	database := "database_handphone"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Brand{}, &models.Color{}, &models.Model{}, &models.Rating{}, &models.Phone{}, &models.Review{}, &models.User{})

	return db
}
