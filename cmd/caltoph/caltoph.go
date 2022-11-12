package main

import (
	"caltoph/internal/api"
	"caltoph/internal/config"
	"caltoph/internal/db"
	"caltoph/internal/health"
	"caltoph/internal/logger"
	"flag"
)

func main() {

	var configFile string
	logger.Init()
	logger.InfoLogger.Println("Starting caltoph")
	flag.StringVar(&configFile, "config", "", "Path to config file")
	flag.Parse()
	serverConfig := config.Init(configFile)
	if serverConfig.DevMode {
		logger.InfoLogger.Println("Caltoph is running in dev mode")
	}
	logger.Init2(serverConfig.Loglevel, serverConfig.DevMode)
	db.Init(serverConfig.Postgres_uri)
	health.Init()
	api.Init()

	select {}
}
