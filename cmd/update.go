package cmd

import (
	"fmt"
	"strconv"
	"todo-go/storage"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing todo item",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("you must specify the ID of the todo item to update and the new task")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID: %v", err)
		}
		newTask := args[1]

		todos, err := storage.LoadTodos()
		if err != nil {
			return err
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Task = newTask
				return storage.SaveTodos(todos)
			}
		}
		return fmt.Errorf("todo item with ID %d not found", id)
	},
}
