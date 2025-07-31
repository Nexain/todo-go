package cmd

import (
	"fmt"
	"strconv"
	"todo-go/storage"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a todo item",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("you must specify the ID of the todo item to delete")
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid ID: %v", err)
		}

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
	},
}
