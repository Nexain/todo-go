package core

import (
	"fmt"
	"todo-go/storage"
)

func UpdateTask(id int, newTask string) (bool, error) {
	todos, err := storage.LoadTodos()
	if err != nil {
		return false, err
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Task = newTask
			return true, storage.SaveTodos(todos)
		}
	}
	return false, fmt.Errorf("todo item with ID %d not found", id)
}
