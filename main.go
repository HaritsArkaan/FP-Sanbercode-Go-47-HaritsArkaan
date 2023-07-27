package main

import (
	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/config"
)

func main() {
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}
