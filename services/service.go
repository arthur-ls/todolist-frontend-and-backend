package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/arthur-ls/todo-list-project/database"
	"github.com/arthur-ls/todo-list-project/entities"
)

var todo entities.TodoItemModel
var db, err = database.InitDB()

func GetAll(w http.ResponseWriter, r *http.Request) {
	var todos []entities.TodoItemModel
	db.Find(&todos)
	jsonString, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonString))
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var todo entities.TodoItemModel
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warn("TodoItem not created")
	}
	json.Unmarshal(reqBody, &todo)
	db.Create(&todo)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"created": true}`)
	json.NewEncoder(w).Encode(todo)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := GetItemByID(id)
	if err == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": false, "error": "Record Not Found"}`)
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Updating TodoItem")
		todo := &entities.TodoItemModel{}
		db.First(&todo, id)
		json.NewDecoder(r.Body).Decode(&todo)
		db.Save(&todo)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"updated": true}`)
	}
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := GetItemByID(id)
	if err == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"get": false, "error": "Record Not Found"}`)
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Getting TodoItem")
		todo := &entities.TodoItemModel{}
		db.First(&todo, id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todo)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := GetItemByID(id)
	if err == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": false, "error": "Record Not Found"}`)
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Deleting TodoItem")
		todo := &entities.TodoItemModel{}
		db.First(&todo, id)
		db.Delete(&todo)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": true}`)
	}
}

func GetItemByID(Id int) bool {
	todo := &entities.TodoItemModel{}
	result := db.First(&todo, Id)
	if result.Error != nil {
		log.Warn("TodoItem not found in database")
		return false
	}
	return true
}
