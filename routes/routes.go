package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
	todos, err := models.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err )
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.Execute(w, todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}

func markTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("Could not parse id", err)
	}

	err = models.MarkDone(id)
	if err != nil {
		fmt.Println("Could not update todo", err)
	}

	sendTodos(w)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Eror parsing form", err)
	}

	err = models.CreateTodo(r.FormValue("todo"))
	if err != nil {
		fmt.Println("Could not create todo", err)
	}

	sendTodos(w)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("Could not parse id", err)
	}

	err = models.DeleteTodo(id)
	if err != nil {
		fmt.Println("Could not delete", err)
	}

	sendTodos(w)
}

func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", mux))
}