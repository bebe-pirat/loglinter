package example

import (
	"log/slog"
	"time"

	"go.uber.org/zap"
)

func correctExamples() {
	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
	slog.Debug("processing request completed")
	slog.Warn("timeout occurred")

	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	sugar.Info("user authenticated successfully")
	sugar.Debug("cache refreshed")
	sugar.Warn("rate limit approaching")

	logger.Info("server started", zap.Int("port", 8080))
	logger.Debug("request processed", zap.Duration("duration", time.Second))
}

func incorrectExamples() {
	slog.Info("Starting server")

	slog.Info("запуск сервера")

	slog.Info("server started!🚀")
	slog.Error("connection failed!!!")

	password := "secret123"
	slog.Info("user password: " + password)

	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	sugar.Info("API Key: abc123")
	sugar.Debug("Token=xyz789")

	apiKey := "key123"
	sugar.Info("api_key=" + apiKey)
}
