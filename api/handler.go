package api

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/prithvipal/todo-app/data"
	"github.com/prithvipal/todo-app/models"
	log "github.com/sirupsen/logrus"
)

var (
	hanlderFunc = map[string]func(w http.ResponseWriter, r *http.Request){
		"GET":    getHanlder,
		"POST":   createHanlder,
		"DELETE": deleteHanlder,
		"PUT":    updateHandler,
		"PATCH":  partialUpdateHandler,
	}
)

type TodoHandler struct {
}

func (th TodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hanlderFunc[r.Method](w, r)
}

func deleteHanlder(w http.ResponseWriter, r *http.Request) {
	id, err := validateUrlAndExtractParam(r.URL.Path)
	if err != nil {
		log.Println("Could not parse request url", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = data.DeleteTodo(id)
	if err != nil {
		log.Println("key not found in database", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}

func createHanlder(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todo.Status != 0 {
		err := fmt.Errorf("status must be NOT_STARTED while creating todo")
		log.Println("error processing request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.SaveTodo(todo)
	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getHanlder(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if url == "/api/v1/todo/" {
		todos := data.ListTodo()
		writeJSON(w, todos)
		return
	}
	id, err := validateUrlAndExtractParam(url)
	if err != nil {
		log.Println("Could not parse request", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	todo, err := data.GetTodo(id)
	if err != nil {
		log.Println("Could not parse request", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, todo)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("In updateHandler  ")
	id, err := validateUrlAndExtractParam(r.URL.Path)
	if err != nil {
		log.Println("Could not parse request url", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.Id = id
	err = data.UpdateTodo(todo)

	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func partialUpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("In partialUpdateHandler  ")
	id, err := validateUrlAndExtractParam(r.URL.Path)
	if err != nil {
		log.Println("Could not parse request url", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	oldTodo, _ := data.GetTodo(id)
	oldTodo.Status = todo.Status
	if todo.Description != "" {
		oldTodo.Description = todo.Description
	}
	fmt.Println(todo.Title)
	if todo.Title != "" {
		oldTodo.Title = todo.Title
	}
	err = data.UpdateTodo(oldTodo)
	if err != nil {
		log.Println("Internal error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
