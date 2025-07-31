package cmd

import (
	"fmt"
	"todo-go/storage"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	RunE: func(cmd *cobra.Command, args []string) error {
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
	},
}
