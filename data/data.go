package data

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/prithvipal/todo-app/models"
)

var (
	database map[string]models.Todo = make(map[string]models.Todo)
)

func SaveTodo(todo models.Todo) {
	todo.Id = uuid.New().String()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	database[todo.Id] = todo
}

func ListTodo(inputs map[string]any) []models.Todo {
	todos := make([]models.Todo, 0)
	title := inputs["title"]
	s := inputs["status"]
	status := s.(*models.StatusType)
	for _, todo := range database {
		if title == "" && status == nil {
			todos = append(todos, todo)
			continue
		}
		if title != "" && status != nil {
			titleStr := title.(string)
			if strings.Contains(todo.Title, string(titleStr)) && todo.Status == *status {
				todos = append(todos, todo)
			}
		} else if title != "" {
			titleStr := title.(string)
			if strings.Contains(todo.Title, string(titleStr)) {
				todos = append(todos, todo)
			}

		} else if todo.Status == *status {
			todos = append(todos, todo)
		}

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
