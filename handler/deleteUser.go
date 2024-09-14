package handler

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api.git/config"
	"github.com/gabrielfmcoelho/abare-api.git/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var user schemas.User

	if err := config.GetSQLite().First(&user, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := config.GetSQLite().Delete(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
