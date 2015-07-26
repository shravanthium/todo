package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var todo Todos

func init() {
	createTodo(Todo{Name: "Write go RESTful API's"})
	createTodo(Todo{Name: "Build Frontend using Angular"})
}

func createTodo(t Todo) Todo {
	todo = append(todo, t)
	return t
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Amigos")
}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if err := json.NewEncoder(w).Encode(todo[id]); err != nil {
		panic(err)
	}

}
