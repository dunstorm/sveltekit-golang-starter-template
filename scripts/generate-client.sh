#!/bin/bash

# Exit on error
set -e

# Navigate to backend directory and generate Swagger docs
cd backend
swag init

# Copy the generated swagger.yaml to frontend
cp docs/swagger.yaml ../frontend/src/lib/swagger.yaml

# Generate TypeScript client using openapi-generator-cli
bunx @openapitools/openapi-generator-cli generate \
  -i docs/swagger.yaml \
  -g typescript-axios \
  -o ../frontend/src/lib/api \
  --additional-properties=supportsES6=true,npmName=@api/client,npmVersion=1.0.0

echo "API client generation completed successfully!"
