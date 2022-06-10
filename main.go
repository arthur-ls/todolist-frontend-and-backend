package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/arthur-ls/todo-list-project/database"
	"github.com/arthur-ls/todo-list-project/migration"
	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, "Hello World")
}

func main() {
	database.InitDB()
	migration.Migrate()
	router := mux.NewRouter()
	router.HandleFunc("/", Handler)
	log.Println(fmt.Sprintf("Starting Server on port 3000"))
	http.ListenAndServe(":3000", router)
}
