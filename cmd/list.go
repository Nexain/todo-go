package cmd

import (
	"fmt"
	"todo-go/core"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo items",
	RunE: func(cmd *cobra.Command, args []string) error {
		todos, err := core.ListTask()
		if err != nil {
			return fmt.Errorf("error listing tasks: %w", err)
		}
		if len(todos) == 0 {
			fmt.Println("No todo items found.")
			return nil
		}
		for _, todo := range todos {
			check := "[ ]"
			if todo.Completed {
				check = "[x]"
			}
			fmt.Printf("%d. %s %s (Created: %s", todo.ID, check, todo.Task, todo.CreatedAt.Format("2006-01-02 15:04"))
			if todo.Completed {
				fmt.Printf(", Completed: %s", todo.CompletedAt.Format("2006-01-02 15:04"))
			}
			fmt.Println(")")
		}
		return nil
	},
}
