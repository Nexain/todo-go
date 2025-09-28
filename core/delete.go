package core

import (
	"context"
	"time"
	"todo-go/db"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteTask(id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	_, err := db.TodoCollection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}
