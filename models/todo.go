package models

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type TodoData struct {
	Todos  []Todo `json:"todos"`
	NextID int    `json:"next_id"`
}
