package core

import (
	"context"
	"time"
	"todo-go/db"
	"todo-go/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ListTask() ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find all todos
	cursor, err := db.TodoCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode into slice of todos
	var todos []models.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	todos = SortTodos(todos) // Sort the todos before returning
	return todos, nil
}
