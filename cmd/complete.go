package cmd

import (
	"fmt"
	"strconv"
	"todo-go/core"

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

		err = core.CompleteTask(id)
		if err != nil {
			return fmt.Errorf("error completing todo item: %v", err)
		}
		return nil
	},
}
