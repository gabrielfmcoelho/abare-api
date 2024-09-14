package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()
	initializeRoutes(router)
	router.Run(":8000")
}
