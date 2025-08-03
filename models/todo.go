package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Task        string    `json:"task"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CompletedAt time.Time `json:"completed_at,omitempty"` // Nullable field for completed tasks
}

type TodoData struct {
	Todos  []Todo `json:"todos"`
	NextID int    `json:"next_id"`
}
