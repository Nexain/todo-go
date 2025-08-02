package core

import (
	"fmt"
	"todo-go/storage"
)

func CompleteTask(id int) error {
	todos, err := storage.LoadTodos()
	if err != nil {
		return err
	}

	for i, todo := range todos {
		if todo.ID == id {
			if todo.Completed {
				return fmt.Errorf("todo item with ID %d is already completed", id)
			}
			todos[i].Completed = true
			return storage.SaveTodos(todos)
		}
	}
	return fmt.Errorf("todo item with ID %d not found", id)
}
