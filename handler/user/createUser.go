package user

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api/handler"
	"github.com/gabrielfmcoelho/abare-api/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api/v1

// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param request body CreateOrUpdateUserRequest true "Request body"
// @Success 200 {object} HandlerUserResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /user [post]
func CreateUserHandler(ctx *gin.Context) {
	userRequest := CreateOrUpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		handler.Logger.Errorf("Failed to bind request: %v", err)
		handler.SendError(ctx, http.StatusBadRequest, "Failed to bind request")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		handler.Logger.Errorf("Failed to hash password: %v", err)
		handler.SendError(ctx, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	userRequest.Password = string(hashedPassword)

	user := schemas.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Gender:    userRequest.Gender,
		Phone:     userRequest.Phone,
		Filiation: userRequest.Filiation,
		Role:      userRequest.Role,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
	}

	if err := handler.Db.Create(&user).Error; err != nil {
		handler.Logger.Errorf("Failed to create user: %v", err)
		handler.SendError(ctx, http.StatusInternalServerError, "Failed to create user")
		return
	}

	handler.SendSuccess(ctx, "User created successfully", user)
}
