package router

import (
	"github.com/gabrielfmcoelho/abare-api/handler"
	"github.com/gabrielfmcoelho/abare-api/handler/user"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	// Set up routes
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/users", user.GetUserHandler)
		v1.GET("/user/:id", user.GetUserHandler)
		v1.POST("/user", user.CreateUserHandler)
		v1.PUT("/user/:id", user.UpdateUserHandler)
		v1.DELETE("/user/:id", user.DeleteUserHandler)
		v1.POST("/login", user.LoginHandler)
	}
}
