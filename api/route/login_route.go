package route

import (
	"time"

	"github.com/gabrielfmcoelho/abare-api/api/controller"
	"github.com/gabrielfmcoelho/abare-api/bootstrap"
	"github.com/gabrielfmcoelho/abare-api/repository"
	"github.com/gabrielfmcoelho/abare-api/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
