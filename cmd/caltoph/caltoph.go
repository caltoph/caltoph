package main

import (
	"caltoph/internal/db"
	"caltoph/internal/logger"
)

func main() {
	logger.Init()
	db.Init()
}
