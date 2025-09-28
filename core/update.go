package core

import (
	"context"
	"time"
	"todo-go/db"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateTask(id int64, newTask string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id, "completed": false}
	update := bson.M{
		"$set": bson.M{
			"task":       newTask,
			"updated_at": time.Now()},
	}
	err := db.TodoCollection.FindOneAndUpdate(ctx, filter, update).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}
