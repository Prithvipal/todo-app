package data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type statusType int

const (
	none statusType = iota
	inProgress
	completed

	max
)

var (
	database map[string]Todo = make(map[string]Todo)
)

type Todo struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      statusType `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func SaveTodo(todo Todo) error {
	todo.Id = uuid.New().String()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	database[todo.Id] = todo
	return nil
}

func ListTodo() []Todo {
	todos := make([]Todo, 0)
	for _, todo := range database {
		todos = append(todos, todo)
	}
	return todos
}

func DeleteTodo(id string) error {
	if _, ok := database[id]; !ok {
		return fmt.Errorf("id not found %v", id)
	}
	delete(database, id)
	return nil
}
