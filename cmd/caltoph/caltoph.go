package main

import (
	"caltoph/internal/db"
	"caltoph/internal/health"
	"caltoph/internal/logger"
)

func main() {
	logger.Init()
	db.Init()
	health.Init()

	select {}
}
