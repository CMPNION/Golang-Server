# Makefile для Go REST API

BINARY_NAME=myapp
MAIN_PACKAGE=cmd/app/main.go

.PHONY: build run clean test lint fmt docker-up

## Сборка бинарника
build:
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)

## Запуск без сборки
run:
	go run $(MAIN_PACKAGE)

## Очистка
clean:
	go clean
	rm -f $(BINARY_NAME)

## Тесты
test:
	go test ./... -v

## Линтинг (golangci-lint должен быть установлен)
lint:
	golangci-lint run

## Форматирование
fmt:
	go fmt ./...

## Docker-compose up (если есть docker-compose.yml)
docker-up:
	docker-compose up --build

## Помощь
help:
	@echo "Доступные команды:"
	@echo "  build       — собрать бинарник"
	@echo "  run         — запустить проект"
	@echo "  clean       — очистить"
	@echo "  test        — запустить тесты"
	@echo "  lint        — запустить линтер"
	@echo "  fmt         — отформатировать код"
	@echo "  docker-up   — запустить docker-compose"
