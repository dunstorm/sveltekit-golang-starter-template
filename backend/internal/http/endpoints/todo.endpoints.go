package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/dunstorm/sveltekit-golang-starter-template/backend/internal/models"
	"github.com/dunstorm/sveltekit-golang-starter-template/backend/internal/repository"
	"github.com/go-chi/chi/v5"
)

// TodoHandler handles HTTP requests for todos
type TodoHandler struct {
	repo *repository.TodoRepository
}

// NewTodoHandler creates a new TodoHandler
func NewTodoHandler(repo *repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

// GetTodos godoc
// @Summary Get all todos
// @Description Get a list of all todo items
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} models.Todo
// @Router /api/todos [get]
func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.repo.GetTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo item with the provided title
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo object"
// @Success 201 {object} models.Todo
// @Router /api/todos [post]
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo := h.repo.CreateTodo(todo.Title)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update a todo's completion status
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body models.Todo true "Todo object"
// @Success 200 {object} models.Todo
// @Router /api/todos/{id} [put]
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var todo struct {
		Completed bool `json:"completed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTodo, found := h.repo.UpdateTodo(id, todo.Completed)
	if !found {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}
