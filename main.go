package main

import (
	"github.com/gabrielfmcoelho/abare-api/config"
	"github.com/gabrielfmcoelho/abare-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization failed: %v", err)
		return
	}
	router.Init()
}
