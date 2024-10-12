package user

import (
	"net/http"

	"github.com/gabrielfmcoelho/abare-api/handler"
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
// @Router /user [post]
func CreateUserHandler(ctx *gin.Context) {
	userRequest := CreateOrUpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	userRequest.Password = string(hashedPassword)
	if err := handler.Db.Create(&userRequest).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
