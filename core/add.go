package core

import (
	"context"
	"fmt"
	"strings"
	"time"
	"todo-go/db"
	"todo-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddTask(task string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the next ID
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.M{"_id": "todo_conter"}
	update := bson.M{"$inc": bson.M{"next_id": 1}}

	var counter struct {
		ID     string `bson:"_id"`
		NextID int64  `bson:"next_id"`
	}

	err := db.DB.Collection("counters").FindOneAndUpdate(ctx,
		filter, update, opts).Decode(&counter)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf("could not get next ID: %v", err)
	}

	// Create new todo
	newTodo := models.Todo{
		ID:        counter.NextID,
		Task:      strings.TrimSpace(task),
		CreatedAt: time.Now(),
		Completed: false,
	}

	// Insert todo to MongoDB
	_, err = db.TodoCollection.InsertOne(ctx, newTodo)
	if err != nil {
		return fmt.Errorf("could not insert todo: %v", err)
	}

	return nil
}
