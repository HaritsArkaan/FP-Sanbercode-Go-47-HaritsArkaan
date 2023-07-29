package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/controller"
	"FP-Sanbercode-Go-47-HaritsArkaanPutranto/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/change-password", controller.ChangePassword)

	r.GET("/brands", controller.GetAllBrand)
	r.GET("/brands/:id", controller.GetBrandById)
	r.GET("/brands/:id/phones", controller.GetPhonesByBrandId)
	r.GET("/brands/:id/models", controller.GetModelsByBrandId)

	brandsMiddlewareRoute := r.Group("/brands")
	brandsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	brandsMiddlewareRoute.POST("/", controller.CreateBrand)
	brandsMiddlewareRoute.PATCH("/:id", controller.UpdateBrand)
	brandsMiddlewareRoute.DELETE(":id", controller.DeleteBrand)

	r.GET("/colors", controller.GetAllColor)
	r.GET("/colors/:id", controller.GetColorById)
	r.GET("/colors/:id/phones", controller.GetPhonesByColorId)

	colorsMiddlewareRoute := r.Group("/colors")
	colorsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	colorsMiddlewareRoute.POST("/", controller.CreateColor)
	colorsMiddlewareRoute.PATCH("/:id", controller.UpdateColor)
	colorsMiddlewareRoute.DELETE("/:id", controller.DeleteColor)

	r.GET("/models", controller.GetAllModel)
	r.GET("/models/:id", controller.GetModelById)
	r.GET("/models/:id/phones", controller.GetPhonesByModelId)

	modelsMiddlewareRoute := r.Group("/models")
	modelsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	modelsMiddlewareRoute.POST("/", controller.CreateModel)
	modelsMiddlewareRoute.PATCH("/:id", controller.UpdateModel)
	modelsMiddlewareRoute.DELETE("/:id", controller.DeleteModel)

	r.GET("/phones", controller.GetAllPhone)
	r.GET("/phones/:id", controller.GetPhoneById)

	phonesMiddlewareRoute := r.Group("/phones")
	phonesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	phonesMiddlewareRoute.POST("/", controller.CreatePhone)
	phonesMiddlewareRoute.PATCH("/:id", controller.UpdatePhone)
	phonesMiddlewareRoute.DELETE("/:id", controller.DeletePhone)

	r.GET("/reviews", controller.GetAllReview)
	r.GET("/reviews/:id", controller.GetReviewById)

	reviewsMiddlewareRoute := r.Group("/reviews")
	reviewsMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	reviewsMiddlewareRoute.POST("/", controller.CreateReview)
	reviewsMiddlewareRoute.PATCH("/:id", controller.UpdateReview)
	reviewsMiddlewareRoute.DELETE("/:id", controller.DeleteReview)

	r.GET("/ratings", controller.GetAllRating)
	r.GET("/ratings/:id", controller.GetRatingById)

	ratingMiddlewareRoute := r.Group("/ratings")
	ratingMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	ratingMiddlewareRoute.POST("/", controller.CreateRating)
	ratingMiddlewareRoute.PATCH("/:id", controller.UpdateRating)
	ratingMiddlewareRoute.DELETE("/:id", controller.DeleteRating)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
