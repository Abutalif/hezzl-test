package main

import (
	"context"
	"goods/configs"
	"goods/internal/services"
	"goods/internal/transport/rest"
	"goods/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "goods-server", log.LstdFlags)
	cfg, err := configs.GetConfigs(".env")
	if err != nil {
		logger.Fatalln(err)
	}
	postgres, err := postgres.InitPostgre(cfg)
	if err != nil {
		logger.Fatalln(err)
	}
	// init redis
	// init clickhouse
	services := services.InitServices(postgres)
	// init services {rougly 3} (uses redis, postgre, clickhouse)

	restSrv := rest.NewServer(cfg, logger, services)
	go func() {
		if err := restSrv.ListenAndServe(); err != nil {
			logger.Fatalln(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err = restSrv.Shutdown(ctx); err != nil {
		logger.Println("Cannot shutdown gracefully:", err)
	}
	logger.Println("Shutting down gracefully")
}
