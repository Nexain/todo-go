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
	return todos, nil
}
