package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/controller"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/brands", controller.GetAllBrand)
	r.POST("/brands", controller.CreateBrand)
	r.GET("/brands/:id", controller.GetBrandById)
	r.GET("/brands/:id/phones", controller.GetPhonesByBrandId)
	r.GET("/brands/:id/models", controller.GetModelsByBrandId)
	r.PATCH("/brands/:id", controller.UpdateBrand)
	r.DELETE("brands/:id", controller.DeleteBrand)

	r.GET("/colors", controller.GetAllColor)
	r.POST("/colors", controller.CreateColor)
	r.GET("/colors/:id", controller.GetColorById)
	r.GET("/colors/:id/phones", controller.GetPhonesByColorId)
	r.PATCH("/colors/:id", controller.UpdateColor)
	r.DELETE("colors/:id", controller.DeleteColor)

	r.GET("/models", controller.GetAllModel)
	r.POST("/models", controller.CreateModel)
	r.GET("/models/:id", controller.GetModelById)
	r.GET("/models/:id/phones", controller.GetPhonesByModelId)
	r.PATCH("/models/:id", controller.UpdateModel)
	r.DELETE("models/:id", controller.DeleteModel)

	r.GET("/phones", controller.GetAllPhone)
	r.POST("/phones", controller.CreatePhone)
	r.GET("/phones/:id", controller.GetPhoneById)
	r.PATCH("/phones/:id", controller.UpdatePhone)
	r.DELETE("phones/:id", controller.DeletePhone)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
