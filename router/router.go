package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	initializeRoutes(router)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	router.Run("0.0.0.0:" + port)
}
