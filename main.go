package main

import (
	"net/http"

	"github.com/Prithvipal/todo-app/api"
)

func main() {
	http.Handle("/api/v1/todo", api.TodoHandler{})
	http.ListenAndServe(":8080", nil)
}
