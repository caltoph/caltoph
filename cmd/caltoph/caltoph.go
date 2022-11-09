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
	flag.StringVar(&configFile, "config", "", "Path to config file")
	flag.Parse()
	serverConfig := config.Init(configFile)
	logger.Init2(serverConfig.Loglevel)
	db.Init(serverConfig.Postgres_uri)
	health.Init()
	api.Init()

	select {}
}
