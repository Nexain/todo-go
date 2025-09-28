package core

import (
	"context"
	"fmt"
	"time"
	"todo-go/db"

	"go.mongodb.org/mongo-driver/bson"
)

func CompleteTask(id int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id, "completed": false}
	update := bson.M{
		"$set": bson.M{"completed": true,
			"completed_at": time.Now(),
			"updated_at":   time.Now()},
	}
	err := db.TodoCollection.FindOneAndUpdate(ctx, filter, update).Err()
	if err != nil {
		return error.Error(err), err
	}

	return fmt.Sprintf("task id: %d marked as completed", id), nil
}
