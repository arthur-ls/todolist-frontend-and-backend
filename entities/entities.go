package entities

type TodoItemModel struct {
	Id          int `gorm:"primary_key"`
	Description string
	Status      string
}
