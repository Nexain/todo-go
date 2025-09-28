package cmd

import (
	"fmt"
	"strconv"
	"todo-go/core"

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

		_, err = core.DeleteTask(int64(id))
		if err != nil {
			return fmt.Errorf("error deleting todo item: %v", err)
		}

		return nil
	},
}
