package router

import (
	"github.com/gabrielfmcoelho/abare-api/handler"
	"github.com/gin-gonic/gin"

	docs "github.com/gabrielfmcoelho/abare-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitHandler()
	// Set up routes
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/users", handler.GetUsersHandler)
		v1.GET("/user/:id", handler.GetUserHandler)
		v1.POST("/user", handler.CreateUserHandler)
		v1.PUT("/user/:id", handler.UpdateUserHandler)
		v1.DELETE("/user/:id", handler.DeleteUserHandler)
		v1.POST("/login", handler.LoginHandler)
	}
	// Initialize Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
