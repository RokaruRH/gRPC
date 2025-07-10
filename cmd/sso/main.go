package main

import (
	"fmt"
	"gRPC/gRPC/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config.MustLoad()

	log := setupLogger(config.Env)

	fmt.Println(config)
	log.Info("START COOKING JESSY!",
		slog.String("env", config.Env),
		slog.Int("env", config.GRPC.Port),
		slog.Any("env", config),
	)

	log.Debug("debug message")
	log.Error("error message")
	log.Warn("warn messahe")

	// Todo Logger
	// TODO app
	// ToDO grpc server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
