package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arthur-ls/todo-list-project/database"
	"github.com/arthur-ls/todo-list-project/entities"
	"github.com/arthur-ls/todo-list-project/migration"
	route "github.com/arthur-ls/todo-list-project/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var todo entities.TodoItemModel

func main() {
	database.InitDB()
	migration.CreateTable()
	router := mux.NewRouter()
	route.Route(router)

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	log.Println(fmt.Sprintf("Starting Server on port 3000"))
	http.ListenAndServe(":3000", handler)
}
