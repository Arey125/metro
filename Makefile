include .env
BIN = metro
NOW = $(shell date '+%Y_%m_%d_%H_%M_%S')

.PHONY: tailwind all run sqlite migrate migrate-down

tailwind:
	@$(TAILWIND) -i app.css -o ./static/tailwind-output.css --minify

all:
	@$(TAILWIND) -i app.css -o ./static/tailwind-output.css --minify
	@templ generate
	@go build -o bin/$(BIN) cmd/web/*

run: all
	@./bin/$(BIN)

sqlite:
	@sqlite3 $(DB)

migrate:
	migrate -database "sqlite3://$(DB)" -path ./migrations up

migrate-down:
	migrate -database "sqlite3://$(DB)" -path ./migrations down 1

backup:
	sqlite3 $(DB) ".backup $(DB_BACKUP_DIR)/$(NOW).db"
