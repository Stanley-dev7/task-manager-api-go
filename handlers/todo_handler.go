package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-learning/day31-task-manager-v2/models"
	"go-learning/day31-task-manager-v2/storage"
)

type Handler struct {
	Todos  []models.Todo
	NextID int
}

// GET + POST /todos
func (h *Handler) TodosHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		WriteJSON(w, 200, true, "Todos fetched successfully", h.Todos)

	case "POST":

		var todo models.Todo

		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil || todo.Title == "" {
			WriteJSON(w, 400, false, "Invalid request body", nil)
			return
		}

		todo.ID = h.NextID
		h.NextID++
		todo.Completed = false

		h.Todos = append(h.Todos, todo)

		storage.SaveTodos(h.Todos)

		WriteJSON(w, 201, true, "Todo created successfully", todo)

	default:
		WriteJSON(w, 405, false, "Method not allowed", nil)
	}
}

// GET + PUT + DELETE /todos/{id}
func (h *Handler) TodoHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		WriteJSON(w, 400, false, "Invalid URL", nil)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		WriteJSON(w, 400, false, "Invalid ID", nil)
		return
	}

	switch r.Method {

	case "GET":
		for _, t := range h.Todos {
			if t.ID == id {
				WriteJSON(w, 200, true, "Todo found", t)
				return
			}
		}
		WriteJSON(w, 404, false, "Todo not found", nil)

	case "PUT":

		var updated models.Todo

		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			WriteJSON(w, 400, false, "Invalid request body", nil)
			return
		}

		for i, t := range h.Todos {
			if t.ID == id {

				h.Todos[i].Title = updated.Title
				h.Todos[i].Completed = updated.Completed

				storage.SaveTodos(h.Todos)

				WriteJSON(w, 200, true, "Todo updated successfully", h.Todos[i])
				return
			}
		}

		WriteJSON(w, 404, false, "Todo not found", nil)

	case "DELETE":

		for i, t := range h.Todos {
			if t.ID == id {

				h.Todos = append(h.Todos[:i], h.Todos[i+1:]...)

				storage.SaveTodos(h.Todos)

				WriteJSON(w, 200, true, "Todo deleted successfully", nil)
				return
			}
		}

		WriteJSON(w, 404, false, "Todo not found", nil)

	default:
		WriteJSON(w, 405, false, "Method not allowed", nil)
	}
}
		