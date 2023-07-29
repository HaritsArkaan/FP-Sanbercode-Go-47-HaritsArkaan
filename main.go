package main

import (
	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/config"
	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/docs"
	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/routes"
	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/utils"
	"log"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	// for load godotenv
	// for env
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Phone Review."
	docs.SwaggerInfo.Version = "1.0"
	envHost := utils.Getenv("ENV_HOST", "localhost")
	envScheme := utils.Getenv("ENV_SCHEME", "http")
	docs.SwaggerInfo.Host = envHost
	docs.SwaggerInfo.Schemes = []string{envScheme}

	//database connection
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run(":" + utils.Getenv("PORT", "8080"))
}
