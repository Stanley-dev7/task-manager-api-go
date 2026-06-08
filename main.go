package main

import (
	"fmt"
	"net/http"

	"go-learning/day31-task-manager-v2/handlers"
	"go-learning/day31-task-manager-v2/models"
	"go-learning/day31-task-manager-v2/storage"
)

func main() {

	// Load saved todos
	todos, _ := storage.LoadTodos()

	// Initialize handler
	h := &handlers.Handler{
		Todos:  todos,
		NextID: getNextID(todos),
	}

	// Routes
	http.HandleFunc("/todos", h.TodosHandler)
	http.HandleFunc("/todos/", h.TodoHandler)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

// Helper function to get next ID
func getNextID(todos []models.Todo) int {
	max := 1
	for _, t := range todos {
		if t.ID >= max {
			max = t.ID + 1
		}
	}
	return max
}