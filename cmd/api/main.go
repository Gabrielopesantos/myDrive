package main

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/utils"
	"log"
	"os"

	postgres "github.com/gabrielopesantos/myDrive-api/pkg/utl/database"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/server"
)

func main() {
	log.Println("Starting API server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Load config %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	psqlDB, err := postgres.NewPsqlDB()
	if err != nil {
		appLogger.Fatalf("Postgres init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, status %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()

	srv := server.NewServer(cfg, psqlDB, appLogger)
	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
