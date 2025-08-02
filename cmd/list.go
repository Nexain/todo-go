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
		err := core.ListTask()
		if err != nil {
			return fmt.Errorf("error listing tasks: %w", err)
		}
		return nil
	},
}
