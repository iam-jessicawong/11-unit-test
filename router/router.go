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

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:id", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.GET("/", controllers.GetProducts)
		productRouter.GET("/:id", middlewares.ProductAuthorization(), controllers.GetProduct)
		productRouter.DELETE("/:id", middlewares.DeleteAuthorization(), middlewares.ProductAuthorization(), controllers.DeleteProduct)
	}

	return r
}
