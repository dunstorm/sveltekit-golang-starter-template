# My Svelte Go Template

This project is a starter template integrating a SvelteKit frontend with a Go backend. The frontend consumes a type-safe API client generated from the backend's Swagger spec, and the backend automatically generates its Swagger documentation using godoc comments and swaggo.

## Project Structure

```
my-svelte-go-template/
├── frontend/
│   ├── src/
│   │   ├── lib/
│   │   │   ├── api/
│   │   │   │   └── client.ts     # Auto-generated API client placeholder
│   │   │   └── swagger.yaml       # Swagger specification reference
│   │   ├── routes/
│   │   │   └── +page.svelte       # Example page using the API client
│   │   └── app.html               # SvelteKit HTML template
│   ├── svelte.config.js           # SvelteKit configuration
│   ├── vite.config.ts             # Vite configuration for SvelteKit
│   ├── package.json               # Frontend dependencies and scripts
│   └── tsconfig.json              # TypeScript configuration
├── backend/
│   ├── api/
│   │   └── handler.go             # API handlers with Swagger annotations
│   ├── docs/
│   │   ├── swagger.json           # Auto-generated Swagger JSON
│   │   └── swagger.yaml           # Auto-generated Swagger YAML
│   ├── main.go                    # Main Go server entry point
│   └── go.mod                     # Go module definition
├── scripts/
│   └── generate-client.sh         # Script to generate the SvelteKit API client from Swagger
├── .github/
│   └── workflows/
│       ├── frontend.yml           # GitHub Actions workflow for the frontend
│       └── backend.yml            # GitHub Actions workflow for the backend
├── .gitignore                   # Git ignore list
└── README.md                    # Project setup instructions
```

## CI/CD with GitHub Actions

- **Frontend**: The GitHub Actions workflow in `.github/workflows/frontend.yml` handles dependency installation, linting, testing, and building the SvelteKit app.
- **Backend**: The GitHub Actions workflow in `.github/workflows/backend.yml` sets up the Go environment, runs tests, generates Swagger documentation, and builds the Go application.
- **Integration**: The CI/CD pipelines include a step to run the client generation script to keep the frontend API client in sync with the backend.

## Setup Instructions

### Prerequisites

- [Node.js](https://nodejs.org/) and [Bun](https://bun.sh/) for frontend development
- [Go](https://golang.org/) (1.18 or newer) for backend development
- [Swaggo](https://github.com/swaggo/swag) for Swagger generation

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/sveltekit-golang-starter-template.git
   cd sveltekit-golang-starter-template
   ```

2. Set up the frontend:
   ```bash
   cd frontend
   bun install
   ```

3. Create a new SvelteKit project using bunx:
   ```bash
   bunx create-svelte@latest .
   ```
   Follow the prompts to configure your SvelteKit project.

4. Set up the backend:
   ```bash
   cd ../backend
   go mod download
   ```

5. Generate the Swagger documentation:
   ```bash
   cd ../backend
   swag init -g main.go -o ./docs
   ```

6. Generate the API client for the frontend:
   ```bash
   cd ..
   ./scripts/generate-client.sh
   ```

### Development

1. Start the frontend development server:
   ```bash
   cd frontend
   bun run dev
   ```

2. Start the backend server:
   ```bash
   cd backend
   go run main.go
   ```

3. Access the application at http://localhost:5173 (frontend) and http://localhost:8080 (backend API)

### Building for Production

1. Build the frontend:
   ```bash
   cd frontend
   bun run build
   ```

2. Build the backend:
   ```bash
   cd backend
   go build -o server main.go
   ```

## License

[Your License Information]
