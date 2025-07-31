package cmd

import (
	"fmt"
	"strconv"
	"todo-go/storage"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a todo item",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("you must specify the ID of the todo item to complete")
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
				todos[i].Completed = true
				return storage.SaveTodos(todos)
			}
		}
		return fmt.Errorf("todo item with ID %d not found", id)
	},
}
