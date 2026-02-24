package testdata

import (
	"go.uber.org/zap"
)

func Zap_log() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Неправильно
	logger.Info("Starting server on port 8080")   // want "first letter in log message is not lower case"
	logger.Error("Failed to connect to database") // want "first letter in log message is not lower case"
	// Правильно
	logger.Info("starting server on port 8080")
	logger.Error("failed to connect to database")

	// Неправильно
	logger.Info("запуск сервера")                    // want "log message contains non-english characters"
	logger.Error("ошибка подключения к базе данных") // want "log message contains non-english characters"
	// Правильно
	logger.Info("starting server")
	logger.Error("failed to connect to database")

	// Неправильно
	logger.Info("server started!🚀")                 // want "log message contains special symbols or emojii"
	logger.Error("connection failed!!!")            // want "log message contains special symbols or emojii"
	logger.Warn("warning: something went wrong...") // want "log message contains special symbols or emojii"
	// Правильно
	logger.Info("server started")
	logger.Error("connection failed")
	logger.Warn("something went wrong")

	password := "1234"
	apiKey := "1234"
	token := "1234"

	// Неправильно
	logger.Info("user password: " + password) // want "log message may contain sensetive or dynamic data"
	logger.Debug("api_key=" + apiKey)         // want "log message may contain sensetive or dynamic data"
	logger.Info("token: " + token)            // want "log message may contain sensetive or dynamic data"

	// Правильно
	logger.Info("user authenticated successfully")
	logger.Debug("api request completed")
	logger.Info("token validated")
}
