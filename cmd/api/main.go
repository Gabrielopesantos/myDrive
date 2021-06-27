package main

import (
	postgres "github.com/gabrielopesantos/myDrive-api/pkg/database/postgres"
	redis "github.com/gabrielopesantos/myDrive-api/pkg/database/redis"
	utils "github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/gabrielopesantos/myDrive-api/config"
	server "github.com/gabrielopesantos/myDrive-api/internal/server"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
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

	// Redis
	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	appLogger.Info("Redis connected")

	// Jaeger
	jaegerCfgInstance := jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           cfg.Jaeger.LogSpans,
			LocalAgentHostPort: cfg.Jaeger.Host,
		},
	}

	tracer, closer, err := jaegerCfgInstance.NewTracer(
		jaegercfg.Logger(jaegerlog.StdLogger),
		jaegercfg.Metrics(metrics.NullFactory),
	)

	if err != nil {
		log.Fatal("cannot create tracer", err)
	}
	appLogger.Info("Jaeger connected")

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	appLogger.Info("Opentracing connected")

	srv := server.NewServer(cfg, psqlDB, redisClient, appLogger)
	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
