package migration

import (
	"github.com/arthur-ls/todo-list-project/database"
	"github.com/arthur-ls/todo-list-project/entities"
)

func CreateTable() {
	db, _ := database.InitDB()
	db.Debug().AutoMigrate(&entities.TodoItemModel{})
}
