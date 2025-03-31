package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TodoEndpoints interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
}

// RegisterTodoRoutes registers todo-related routes
func RegisterTodoRoutes(r chi.Router, e TodoEndpoints) {
	r.Route("/api/todos", func(r chi.Router) {
		r.Get("/", e.GetTodos)
		r.Post("/", e.CreateTodo)
		r.Put("/{id}", e.UpdateTodo)
	})
}
