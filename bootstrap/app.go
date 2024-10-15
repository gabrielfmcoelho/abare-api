package bootstrap

import (
	"log"

	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = NewDatabaseConnection(app.Env)

	// Run auto-migration
	AutoMigrate(app.DB)

	return *app
}

func (app *Application) CloseDBConnection() {
	sqlDB, err := app.DB.DB()
	if err != nil {
		log.Fatal("Failed to get database object from Gorm DB:", err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatal("Failed to close database connection:", err)
	}
}
