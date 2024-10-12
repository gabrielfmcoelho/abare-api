package user

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api/handler"
	"github.com/gabrielfmcoelho/abare-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} HandlerUsersResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /users [get]
func GetUsersHandler(ctx *gin.Context) {
	users := []schemas.User{}

	if err := handler.Db.Find(&users).Error; err != nil {
		handler.Logger.Error("Error fetching users", err)
		handler.SendError(ctx, http.StatusInternalServerError, "Error fetching users")
		return
	}

	handler.SendSuccess(ctx, "Users fetched successfully", users)
}
