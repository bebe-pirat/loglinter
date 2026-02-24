package testdata

import (
	"log/slog"
)

func Slog_log() {
	// Неправильно
	slog.Info("Starting server on port 8080")   // want "first letter in log message is not lower case"
	slog.Error("Failed to connect to database") // want "first letter in log message is not lower case"
	// Правильно
	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")

	// Неправильно
	slog.Info("запуск сервера")                    // want "log message contains non-english characters"
	slog.Error("ошибка подключения к базе данных") // want "log message contains non-english characters"
	// Правильно
	slog.Info("starting server")
	slog.Error("failed to connect to database")

	// Неправильно
	slog.Info("server started!🚀")                 // want "log message contains special symbols or emojii"
	slog.Error("connection failed!!!")            // want "log message contains special symbols or emojii"
	slog.Warn("warning: something went wrong...") // want "log message contains special symbols or emojii"
	// Правильно
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")

	password := "1234"
	apiKey := "1234"
	token := "1234"
	// Неправильно
	slog.Info("user password: " + password) // want "log message may contain sensetive or dynamic data"
	slog.Debug("api_key=" + apiKey)         // want "log message may contain sensetive or dynamic data"
	slog.Info("token: " + token)            // want "log message may contain sensetive or dynamic data"
	// Правильно
	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token validated")

	port := 8080
	slog.Debug("server started", "port", port)
}
