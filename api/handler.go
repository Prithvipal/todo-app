package api

import (
	"encoding/json"
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
	hanlderFunc[r.Method](w, r)
}

func deleteHanlder(w http.ResponseWriter, r *http.Request) {
	logrus.Info("In Delete Handler")
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
	w.Write([]byte("hello"))
}
func partialUpdateHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("In partialUpdateHandler  ")
	w.Write([]byte("hello"))
}

func writeJSON(w http.ResponseWriter, records any) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(records)
	if err != nil {
		log.Println("Error while getting TODO List", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}
