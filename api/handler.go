package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Prithvipal/todo-app/data"
	"github.com/sirupsen/logrus"
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
	fmt.Println("servads", r.Method)
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
	var todo data.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
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
	todos := data.ListTodo()
	writeJSON(w, todos)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("In updateHandler  ")
	id, err := validateUrlAndExtractParam(r.URL.Path)
	if err != nil {
		log.Println("Could not parse request url", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var todo data.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Could not parse request payload", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todo.Status < 0 || todo.Status > data.MaxStatus-1 {
		err := fmt.Errorf("valid range of status is 0 to %v", data.MaxStatus-1)
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
	logrus.Info("In partialUpdateHandler  ")
	w.Write([]byte("hello"))
}
