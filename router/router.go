package router

import (
	"unit-test/controllers"
	"unit-test/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
	}

	productController := controllers.ProductController{}
	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:id", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.GET("/", productController.GetProducts)
		productRouter.GET("/:id", middlewares.ProductAuthorization(), productController.GetProduct)
		productRouter.DELETE("/:id", middlewares.DeleteAuthorization(), middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
