package controller

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gabrielfmcoelho/abare-api/bootstrap"
	"github.com/gabrielfmcoelho/abare-api/domain"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

// User Login
// @Summary Login user
// @Description Authenticates a user using their email and password, then returns access and refresh tokens for session management.
// @Tags User
// @ID login
// @Accept json
// @Produce json
// @Param loginRequest body domain.LoginRequest true "Login Request"
// @Success 200 {object} domain.LoginResponse "Successful login, returns access and refresh tokens"
// @Failure 400 {object} domain.ErrorResponse "Bad Request - Invalid input"
// @Failure 401 {object} domain.ErrorResponse "Unauthorized - Incorrect email or password"
// @Failure 404 {object} domain.ErrorResponse "Not Found - User not found"
// @Failure 500 {object} domain.ErrorResponse "Internal Server Error"
// @Router /login [post]
func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
