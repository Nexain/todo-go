package cmd

import (
	"fmt"
	"strings"
	"todo-go/core"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new todo item",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		task := strings.Join(args, " ")
		err := core.AddTask(task)
		if err != nil {
			return fmt.Errorf("could not add task: %v", err)
		}
		fmt.Printf("✅ Todo added: %s\n", task)
		return nil
	},
}
