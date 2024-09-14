package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("failed to initialize SQLite: %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func SetSQLite(d *gorm.DB) {
	db = d
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
