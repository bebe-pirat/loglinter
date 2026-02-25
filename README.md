# loglinter

Линтер кода для Go, который проверяет сообщения логов на соответствие требований: 
* первая буква сообщения лога должна быть строчной
* сообщения должны содержать только английские символы
* сообщения не должны содержать специальные символы и эмозди
* сообщения не должны содержать чувствительную информацию

## Установка
### Предварительные требования
* Go 1.22+
* golangci-lint версии 2.0 или выше

### Сборка
1. git clone https://github.com/bebe-pirat/loglinter.git
2. cd loglinter
3. golangci-lint custom -v

## Запуск 
1. ./bin/my-golangci-lint run [путь к директории с файлами или самому файлу]

**Пример запуска**
1. golangci-lint custom -v 
2. ./bin/my-golangci-lint run ./example.go

**Пример вывода линтера**
```
example.go:28:12: first letter in log message should be lowercase (loglinter)
        slog.Info("Starting server")
                  ^
example.go:30:12: log message should contain only English characters (loglinter)
        slog.Info("запуск сервера")
                  ^
example.go:32:12: log message contains special symbols or emoji (loglinter)
        slog.Info("server started!🚀")
                  ^
example.go:33:13: log message contains special symbols or emoji (loglinter)
        slog.Error("connection failed!!!")
                   ^
example.go:36:12: log message may contain sensitive data (loglinter)
        slog.Info("user password: " + password)
                  ^
example.go:41:13: first letter in log message should be lowercase (loglinter)
        sugar.Info("API Key: abc123")
                   ^
example.go:42:14: first letter in log message should be lowercase (loglinter)
        sugar.Debug("Token=xyz789")
                    ^
example.go:45:13: log message may contain sensitive data (loglinter)
        sugar.Info("api_key=" + apiKey)
                   ^
8 issues:
loglinter: 8
```

**Запуск тестов**
1. cd pkg/analyzer
2. go test


