package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rpturbina/learn-go-fga/belajar-jwt/controllers"
	"github.com/rpturbina/learn-go-fga/belajar-jwt/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

	return r
}
