APP_NAME := regista-api
MIGRATE ?= migrate
DATABASE_URL ?= $(shell echo $$DATABASE_URL)

.PHONY: run build test migrate-up migrate-down migrate-create docker-up docker-down

run:
	go run ./cmd/api

build:
	go build -o bin/$(APP_NAME) ./cmd/api

test:
	go test ./...

migrate-up:
	$(MIGRATE) -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	$(MIGRATE) -path migrations -database "$(DATABASE_URL)" down 1

migrate-create:
	@test -n "$(NAME)" || (echo "Usage: make migrate-create NAME=description" && exit 1)
	$(MIGRATE) create -ext sql -dir migrations -seq $(NAME)

docker-up:
	docker compose up -d

docker-down:
	docker compose down
