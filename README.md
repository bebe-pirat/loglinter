**Запуск проекта**
1. git clone https://github.com/bebe-pirat/loglinter.git
2. cd loglinter
3. golangci-lint custom -v 
4. ./bin/my-golangci-lint run [путь к директории с файлами или самому файлу]

**Пример запуска**
1. golangci-lint custom -v 
2. ./bin/my-golangci-lint run ./example.go

**Запуск тестов**
1. cd pkg/analyzer
2. go test
