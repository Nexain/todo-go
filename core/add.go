package core

import (
	"fmt"
	"strings"
	"time"
	"todo-go/models"
	"todo-go/storage"
)

func AddTask(task string) error {
	data, err := storage.LoadData()
	if err != nil {
		return fmt.Errorf("could not load data: %v", err)
	}

	if data.NextID < 1 { //making the id start from 1
		data.NextID = 1
	}

	newTodo := models.Todo{
		ID:        data.NextID,
		Task:      strings.TrimSpace(task),
		CreatedAt: time.Now(),
	}
	data.Todos = append(data.Todos, newTodo)
	data.NextID++

	return storage.SaveData(data)
}
