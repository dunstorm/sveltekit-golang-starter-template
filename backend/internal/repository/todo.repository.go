package repository

import "github.com/dunstorm/sveltekit-golang-starter-template/backend/internal/models"

// TodoRepository handles todo storage operations
type TodoRepository struct {
	todos []models.Todo
}

// NewTodoRepository creates a new TodoRepository instance
func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos: []models.Todo{
			{ID: 1, Title: "Learn SvelteKit", Completed: false},
			{ID: 2, Title: "Learn Go", Completed: false},
		},
	}
}

// GetTodos returns all todos
func (r *TodoRepository) GetTodos() []models.Todo {
	return r.todos
}

// CreateTodo adds a new todo to the list
func (r *TodoRepository) CreateTodo(title string) models.Todo {
	todo := models.Todo{
		ID:        len(r.todos) + 1,
		Title:     title,
		Completed: false,
	}
	r.todos = append(r.todos, todo)
	return todo
}

// UpdateTodo updates a todo's status
func (r *TodoRepository) UpdateTodo(id int, completed bool) (models.Todo, bool) {
	for i, todo := range r.todos {
		if todo.ID == id {
			r.todos[i].Completed = completed
			return r.todos[i], true
		}
	}
	return models.Todo{}, false
}
