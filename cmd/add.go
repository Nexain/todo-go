package cmd

import (
	"fmt"
	"strings"
	"todo-go/models"
	"todo-go/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new todo item",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := storage.LoadData()
		if err != nil {
			return fmt.Errorf("could not load data: %v", err)
		}

		newTodo := models.Todo{
			ID:   data.NextID,
			Task: strings.Join(args, " "),
		}
		data.Todos = append(data.Todos, newTodo)
		data.NextID++

		err = storage.SaveData(data)
		if err != nil {
			return fmt.Errorf("could not save todo: %v", err)
		}
		fmt.Printf("✅ Todo added: %s\n", newTodo.Task)
		return nil
	},
}
