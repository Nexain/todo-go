package core

import (
	"fmt"
	"time"
	"todo-go/storage"
)

func CompleteTask(id int) (string, error) {
	todos, err := storage.LoadTodos()
	if err != nil {
		return "err", err
	}

	for i, todo := range todos {
		if todo.ID == id {
			if todo.Completed {
				return "already_done", fmt.Errorf("todo item with ID %d is already completed", id)
			}
			todos[i].Completed = true
			todos[i].CompletedAt = time.Now()
			return "ok", storage.SaveTodos(todos)
		}
	}
	return "not_found", fmt.Errorf("todo item with ID %d not found", id)
}
