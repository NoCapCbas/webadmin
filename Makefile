# Variables
DOCKER_COMPOSE = docker compose

# Phony targets
.PHONY: dev

# Default target
all: dev

# Development
dev:
	@if [ ! -f .env.dev ]; then \
		echo "Error: .env.dev file not found. Please create the file and try again."; \
		exit 1; \
	fi
	set -a; source .env.dev; set +a; \
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml up -d --build --force-recreate

# Help
help:
	@echo "Available targets:"
	@echo "  dev : Run the development environment using Docker Compose"
