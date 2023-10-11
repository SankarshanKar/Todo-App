package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"todo-app/models"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {
	todos, err := models.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from the DB", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	
	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sendTodos(w)
}

func markTodo(w http.ResponseWriter, r *http.Request) {
	
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	
}

func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", mux))
}