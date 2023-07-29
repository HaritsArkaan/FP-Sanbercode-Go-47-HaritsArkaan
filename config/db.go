package config

import (
	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/models"
	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := utils.Getenv("MYSQLUSER", "root")
	password := utils.Getenv("MYSQLPASSWORD", "cACgKjvnypF9zuIOPkxd")
	host := utils.Getenv("MYSQLHOST", "containers-us-west-71.railway.app")
	port := utils.Getenv("MYSQLPORT", "6578")
	database := utils.Getenv("MYSQLDATABASE", "railway")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Brand{}, &models.Color{}, &models.Model{}, &models.Rating{}, &models.Phone{}, &models.Review{}, &models.User{})

	return db
}
