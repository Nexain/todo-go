package models

import (
	"time"
)

type Todo struct {
	ID          int64     `json:"id" bson:"id"`
	Task        string    `json:"task" bson:"task"`
	Completed   bool      `json:"completed" bson:"completed"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	CompletedAt time.Time `json:"completed_at" bson:"completed_at,omitempty"`
}

type TodoData struct {
	Todos  []Todo `json:"todos"`
	NextID int    `json:"next_id"`
}
