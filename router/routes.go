package router

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1") // Similar to apiRouter from fastAPI from Python
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "pong",
			})
		})
		v1.GET("/users", handler.GetUsersHandler)
		v1.GET("/user/:id", handler.GetUserHandler)
		v1.POST("/user", handler.CreateUserHandler)
		v1.PUT("/user/:id", handler.UpdateUserHandler)
		v1.DELETE("/user/:id", handler.DeleteUserHandler)
		v1.POST("/login", handler.LoginHandler)
	}
}
