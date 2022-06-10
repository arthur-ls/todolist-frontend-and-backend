package migration

import (
	"github.com/arthur-ls/todo-list-project/database"
	"github.com/arthur-ls/todo-list-project/entities"
)

func Migrate() {
	db, _ := database.InitDB()
	//db.Debug().DropTableIfExists(&entities.TodoItemModel{})
	db.Debug().AutoMigrate(&entities.TodoItemModel{})
}
