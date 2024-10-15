package route

import (
	"time"

	"github.com/gabrielfmcoelho/abare-api/api/middleware"
	"github.com/gabrielfmcoelho/abare-api/bootstrap"
	_ "github.com/gabrielfmcoelho/abare-api/docs"
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
	_ "github.com/swaggo/swag"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, router *gin.Engine) {
	// Router documentation binding
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./docs/swagger.json",
		SpecPath:    "/docs/swagger.json",
		DocsPath:    "/docs/",
	}
	router.GET("/docs/*any", ginredoc.New(doc))

	// All Public APIs
	publicRouter := router.Group("")
	//NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	//NewRefreshTokenRouter(env, timeout, db, publicRouter)

	// All Private APIs
	protectedRouter := router.Group("")
	/// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	//NewProfileRouter(env, timeout, db, protectedRouter)
	//NewTaskRouter(env, timeout, db, protectedRouter)
}
