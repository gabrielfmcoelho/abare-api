package user

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api/handler"
	"github.com/gabrielfmcoelho/abare-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Get a user
// @Description Get a user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} HandlerUserResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Router /user/{id} [get]
func GetUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handler.Logger.Error("User ID not provided")
		handler.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := schemas.User{}

	if err := handler.Db.First(&user, id).Error; err != nil {
		handler.Logger.Error("User not found", err)
		handler.SendError(ctx, http.StatusNotFound, "User not found")
	}

	handler.SendSuccess(ctx, "User found successfully", user)
}
