.PHONY: local run build test

# ==============================================================================
# Go migrate postgresql
force:
	migrate -database postgresql://gabriel:leirbag123@localhost:5432/users?sslmode=disable -path migrations force 1

version:
	migrate -database postgresql://gabriel:leirbag123@localhost:5432/users?sslmode=disable -path migrations version 1

migrate-up:
	migrate -database postgresql://gabriel:leirbag123@localhost:5432/users?sslmode=disable -path migrations up 1

migrate-down:
	migrate -database postgresql://gabriel:leirbag123@localhost:5432/users?sslmode=disable -path migrations down 1

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
