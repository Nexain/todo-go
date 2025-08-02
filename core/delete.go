package core

import (
	"fmt"
	"todo-go/storage"
)

func DeleteTask(id int) error {
	todos, err := storage.LoadTodos()
	if err != nil {
		return err
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return storage.SaveTodos(todos)
		}
	}
	return fmt.Errorf("todo item with ID %d not found", id)
}
