.PHONY: local run build test


# ==============================================================================
# Docker compose commands


local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build -d


# ==============================================================================
# Tools commands

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

swaggo:
	echo "Starting swagger generating"
	swag init -g **/**/*.go

# ==============================================================================
# Main

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor


# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)
