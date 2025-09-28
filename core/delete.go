package core

import (
	"fmt"
	"todo-go/storage"
)

func DeleteTask(id int64) (bool, error) {
	todos, err := storage.LoadTodos()
	if err != nil {
		return false, err
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true, storage.SaveTodos(todos)
		}
	}
	return false, fmt.Errorf("todo item with ID %d not found", id)
}
