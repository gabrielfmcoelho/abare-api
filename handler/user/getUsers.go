package user

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api/config"
	"github.com/gabrielfmcoelho/abare-api/schemas"
	"github.com/gin-gonic/gin"
)

func GetUsersHandler(ctx *gin.Context) {
	var users []schemas.User

	if err := config.GetSQLite().Find(&users).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
