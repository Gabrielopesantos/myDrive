.PHONY: local build test

# Docker compose commands

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build -d

local-down:
	echo "Stopping local environment"
	docker-compose down

# Main

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...