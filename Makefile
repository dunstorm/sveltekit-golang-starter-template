# Makefile for sveltekit-golang-starter-template

# Variables
BACKEND_DIR = ./backend
FRONTEND_DIR = ./frontend
SWAGGER_OUTPUT = $(BACKEND_DIR)/docs
FRONTEND_API_DIR = $(FRONTEND_DIR)/src/lib/api
FRONTEND_SWAGGER_FILE = $(FRONTEND_DIR)/src/lib/swagger.yaml

.PHONY: all help gen/swagger gen/client gen/all

# Default target
all: gen/all

# Help message
help:
	@echo "Available targets:"
	@echo "  gen/swagger - Generate Swagger documentation from Go code"
	@echo "  gen/client  - Generate TypeScript client from Swagger"
	@echo "  gen/all     - Run all generation steps"
	@echo "  help        - Show this help message"

# Generate Swagger documentation
gen/swagger:
	@echo "Generating Swagger documentation..."
	cd $(BACKEND_DIR) && swag init -g cmd/server/main.go --parseDependency --parseInternal
	@echo "Swagger documentation generated in $(SWAGGER_OUTPUT)"

# Copy Swagger and generate TypeScript client
gen/client: gen/swagger
	@echo "Copying Swagger YAML to frontend..."
	cp $(SWAGGER_OUTPUT)/swagger.yaml $(FRONTEND_SWAGGER_FILE)
	@echo "Generating TypeScript client..."
	cd $(BACKEND_DIR) && bunx @openapitools/openapi-generator-cli generate \
		-i docs/swagger.yaml \
		-g typescript-axios \
		-o ../$(FRONTEND_API_DIR) \
		--additional-properties=supportsES6=true,npmName=@api/client,npmVersion=1.0.0
	@echo "TypeScript client generated in $(FRONTEND_API_DIR)"

# Generate all - swagger + client
gen/all: gen/client
	@echo "All API artifacts generated successfully!"

# Run backend in dev mode
dev/backend:
	cd $(BACKEND_DIR) && go run cmd/server/main.go

# Run frontend in dev mode
dev/frontend:
	cd $(FRONTEND_DIR) && bun run dev
