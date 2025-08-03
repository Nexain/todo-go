package core

import (
	"todo-go/models"
	"todo-go/storage"
)

func ListTask() ([]models.Todo, error) {
	todos, err := storage.LoadTodos()
	if err != nil {
		return []models.Todo{}, err
	}
	todos = SortTodos(todos) // Sort the todos before returning
	return todos, nil
}
