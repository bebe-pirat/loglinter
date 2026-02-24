package p

import (
	"log/slog"
)

func Slog_log() {
	// Неправильно
	slog.Info("Starting server on port 8080")   // want "first letter in log message should be lowercase"
	slog.Error("Failed to connect to database") // want "first letter in log message should be lowercase"
	// Правильно
	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")

	// Неправильно
	slog.Info("запуск сервера")                    // want "log message should contain only English characters"
	slog.Error("ошибка подключения к базе данных") // want "log message should contain only English characters"
	// Правильно
	slog.Info("starting server")
	slog.Error("failed to connect to database")

	// Неправильно
	slog.Info("server started!🚀")                 // want "log message contains special symbols or emoji"
	slog.Error("connection failed!!!")            // want "log message contains special symbols or emoji"
	slog.Warn("warning: something went wrong...") // want "log message contains special symbols or emoji"
	// Правильно
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")

	password := "1234"
	apiKey := "1234"
	token := "1234"
	// Неправильно
	slog.Info("user password: " + password) // want "log message may contain sensitive data"
	slog.Debug("api_key=" + apiKey)         // want "log message may contain sensitive data"
	slog.Info("token: " + token)            // want "log message may contain sensitive data"
	// Правильно
	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token validated")

	port := 8080
	slog.Debug("server started", "port", port)
}
