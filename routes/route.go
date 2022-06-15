package route

import (
	service "github.com/arthur-ls/todo-list-project/services"
	"github.com/gorilla/mux"
)

func Route(router *mux.Router) {
	router.HandleFunc("/todo", service.GetAll).Methods("GET")
	router.HandleFunc("/todo/{id}", service.GetItem).Methods("GET")
	router.HandleFunc("/todo", service.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", service.UpdateItem).Methods("PUT")
	router.HandleFunc("/todo/{id}", service.DeleteItem).Methods("DELETE")
}
