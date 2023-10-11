package main

import (
	"todo-app/models"
	"todo-app/routes"
)

func main() {
	models.Setup()
	routes.SetupAndRun()
}