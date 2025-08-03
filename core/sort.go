package core

import (
	"sort"
	"todo-go/models"
)

func SortTodos(todos []models.Todo) []models.Todo {
	sort.Slice(todos, func(i, j int) bool {
		// 1️⃣ Place uncompleted todos at the top
		if !todos[i].Completed && todos[j].Completed {
			return true
		}
		if todos[i].Completed && !todos[j].Completed {
			return false
		}

		// 2️⃣ For uncompleted todos → sort by CreatedAt DESC
		if !todos[i].Completed && !todos[j].Completed {
			return todos[i].CreatedAt.After(todos[j].CreatedAt)
		}

		// 3️⃣ For completed todos → sort by CompletedAt DESC
		if todos[i].Completed && todos[j].Completed {
			return todos[i].CompletedAt.After(todos[j].CompletedAt)
		}

		return false
	})
	return todos
}
