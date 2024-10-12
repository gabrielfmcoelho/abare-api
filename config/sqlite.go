package config

import (
	"os"

	"github.com/gabrielfmcoelho/abare-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "abare.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("database file does not exist, creating a new one")
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("failed to create database file: %v", err)
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect to database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		logger.Errorf("failed to auto migrate schema: %v", err)
		return nil, err
	}
	return db, nil
}
