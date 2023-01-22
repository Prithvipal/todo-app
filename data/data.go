package data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/prithvipal/todo-app/models"
)

var (
	database map[string]models.Todo = make(map[string]models.Todo)
)

func SaveTodo(todo models.Todo) error {
	todo.Id = uuid.New().String()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	database[todo.Id] = todo
	return nil
}

func ListTodo() []models.Todo {
	todos := make([]models.Todo, 0)
	for _, todo := range database {
		todos = append(todos, todo)
	}
	return todos
}

func GetTodo(id string) (models.Todo, error) {
	todo, ok := database[id]
	if !ok {
		return models.Todo{}, fmt.Errorf("id not found %v", id)
	}
	return todo, nil
}

func DeleteTodo(id string) error {
	if _, ok := database[id]; !ok {
		return fmt.Errorf("id not found %v", id)
	}
	delete(database, id)
	return nil
}

func UpdateTodo(todo models.Todo) error {
	if _, ok := database[todo.Id]; !ok {
		return fmt.Errorf("id not found %v", todo.Id)
	}

	todo.UpdatedAt = time.Now()
	todo.CreatedAt = database[todo.Id].CreatedAt
	database[todo.Id] = todo
	return nil
}
