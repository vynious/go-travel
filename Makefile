


schema-up:
	goose -dir internal/db/migration postgres "user=shawntyw password=shawntyw dbname=godb sslmode=disable" up

schema-down:
	goose -dir internal/db/migration postgres "user=shawntyw password=shawntyw dbname=godb sslmode=disable" down

query:
	sqlc generate

#make run:
#	cd cmd && go run main.go
.PHONY: build run


APP_NAME := go-travel
BUILD_DIR := ./bin

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR) ./cmd

run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/cmd

