package cmd

import (
	"fmt"
	"strconv"
	"todo-go/core"

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

		_, err = core.UpdateTask(id, newTask)
		if err != nil {
			return fmt.Errorf("failed to update task: %v", err)
		}
		fmt.Printf("Todo item with ID %d updated successfully to: %s\n", id, newTask)
		return nil
	},
}
