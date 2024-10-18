package main

import (
	"log"
	"time"

	"github.com/gabrielfmcoelho/abare-api/api/route"
	"github.com/gabrielfmcoelho/abare-api/bootstrap"
	"github.com/gin-gonic/gin"
)

// @title           Abaré API
// @version         0.1.1
// @description		Abaré API is a RESTful API for managing users, diaries, families and children related to follow-up programs for the development of children and teenagers with autism spectrum disorder.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Eng. Gabriel Coelho, Eng. João Vitor, Eng. Samuel Martins | Abaré development team
// @contact.url    https://abare.com.br
// @contact.email  suporte@abare.com.br

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Initialize the application
	app := bootstrap.App()
	defer app.CloseDBConnection()

	// Configuration variables
	env := app.Env

	// Database instance (Gorm DB)
	db := app.DB

	// Context timeout
	timeout := time.Duration(env.ContextTimeout) * time.Second

	// Create a Gin router instance
	router := gin.Default()

	// Route binding
	route.Setup(env, timeout, db, router)

	// Run the server
	if err := router.Run(env.ServerAddress); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
