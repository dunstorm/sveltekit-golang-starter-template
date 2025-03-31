package app

import (
	"log"
	"net/http"
	"time"

	// Import the docs package
	_ "github.com/dunstorm/sveltekit-golang-starter-template/backend/docs"
	"github.com/dunstorm/sveltekit-golang-starter-template/backend/internal/http/endpoints"
	"github.com/dunstorm/sveltekit-golang-starter-template/backend/internal/http/routes"
	"github.com/dunstorm/sveltekit-golang-starter-template/backend/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (e *engine) startWebServer() {
	router := chi.NewRouter()

	// Add basic middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	// Allow CORS for all origins
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler)

	// Serve swagger.json file directly
	router.Get("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	// Serve Swagger documentation
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.json"),
		httpSwagger.DocExpansion("none"),
		httpSwagger.Layout(httpSwagger.BaseLayout),
	))

	endpoints := e.createEndpoints()

	// Register routes
	routes.RegisterTodoRoutes(router, endpoints.Todo)

	// Start the HTTP server in a goroutine
	e.runners.Add(1)
	go func() {
		defer e.runners.Done()

		port := e.config.HTTP.Port
		log.Printf("Starting server on port %s", port)

		server := &http.Server{
			Addr:         ":" + port,
			Handler:      router,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		}

		// Graceful shutdown can be added here
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
}

func (e *engine) createEndpoints() routes.Endpoints {
	todoRepository := repository.NewTodoRepository()
	todoEndpoints := endpoints.NewTodoHandler(todoRepository)

	return routes.Endpoints{
		Todo: todoEndpoints,
	}
}
