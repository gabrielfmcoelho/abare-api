package handler

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api.git/config"
	"github.com/gabrielfmcoelho/abare-api.git/schemas"
	"github.com/gin-gonic/gin"
)

func GetUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var user schemas.User

	if err := config.GetSQLite().First(&user, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
