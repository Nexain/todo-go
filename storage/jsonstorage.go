package storage

import (
	"encoding/json"
	"os"
	"todo-go/models"
)

var filePath = "todos.json"

func LoadData() (models.TodoData, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return models.TodoData{}, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return models.TodoData{}, err
	}

	var todoData models.TodoData
	err = json.Unmarshal(data, &todoData)
	if err != nil {
		return models.TodoData{}, err
	}

	return todoData, nil
}

func SaveData(data models.TodoData) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonData, 0644)
}

func LoadTodos() ([]models.Todo, error) {
	data, err := LoadData()
	if err != nil {
		return []models.Todo{}, err
	}
	return data.Todos, nil
}

func SaveTodos(todos []models.Todo) error {
	data, err := LoadData()
	if err != nil {
		return err
	}

	data.Todos = todos
	return SaveData(data)
}
