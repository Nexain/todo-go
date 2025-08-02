package core

import (
	"fmt"
	"todo-go/storage"
)

func ListTask() error {
	todos, err := storage.LoadTodos()
	if err != nil {
		return err
	}
	if len(todos) == 0 {
		fmt.Println("No todo items found.")
		return nil
	}
	for _, todo := range todos {
		check := "[ ]"
		if todo.Completed {
			check = "[v]"
		}
		fmt.Printf("%d. %s %s\n", todo.ID, check, todo.Task)
	}
	return nil
}
