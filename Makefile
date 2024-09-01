# Define the variables
PROJECT_NAME = user-management
# Default PORT value if not provided
PORT ?= 8080

# Commands to run
.PHONY: help tidy run test coverage docker-build docker-run

help:
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  tidy                    Format the code, vendor, and tidy the Go modules"
	@echo "  run [PORT=9090]         Runs the application on the specified port. If no port is provided, it defaults to $(PORT)."
	@echo "                          Example: make run PORT=9090"
	@echo "  test                    Runs all the tests in the sub folders."
	@echo "  coverage                Runs all the tests cases and finds coverage."
	@echo "  docker-build            Builds the docker image"
	@echo "  docker-run [PORT=9090]  Runs the docker container"

tidy:
	go fmt ./...
	go mod tidy
	go mod vendor

run:
	go run cmd/server/main.go -port=$(PORT)

test:
	go test ./...

coverage:
	go test ./... -coverprofile coverage.out
	go tool cover -func coverage.out

docker-build:
	@echo "Building Docker Image..."
	docker build --rm -t user-management .

docker-run:
	@echo "Building Docker Container..."
	docker run --rm -it -p $(PORT):8080 --name user-managementapp user-management
