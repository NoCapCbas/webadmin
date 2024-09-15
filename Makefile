# Variables
DOCKER_COMPOSE = docker compose

# Phony targets
.PHONY: dev test prod

# Default target
all: dev

# Development
dev:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml up -d --build --force-recreate

# Run tests within development environment
test:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml exec web-admin go test ./...

# Run tests within development environment
test-mongo:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml exec web-admin go test ./data/mongo


# Production
prod:
	$(DOCKER_COMPOSE) -f docker-compose.prod.yml up -d --build --force-recreate

# Help
help:
	@echo "Available targets:"
	@echo "  dev : Run the development environment using Docker Compose"
	@echo "  test : Run the tests"
	@echo "  prod : Run the production environment using Docker Compose"
