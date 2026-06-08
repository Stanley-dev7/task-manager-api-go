package storage

import (
	"encoding/json"
	"os"

	"go-learning/day31-task-manager-v2/models"
)

const fileName = "data/todos.json"

// SaveTodos writes todos into JSON file
func SaveTodos(todos []models.Todo) error {

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(todos)
}

// LoadTodos reads todos from JSON file
func LoadTodos() ([]models.Todo, error) {

	var todos []models.Todo

	file, err := os.Open(fileName)
	if err != nil {
		// if file doesn't exist, return empty slice
		return todos, nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&todos)
	if err != nil {
		return []models.Todo{}, nil
	}

	return todos, nil
}