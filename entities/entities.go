package entities

type TodoItemModel struct {
	Id          int    `gorm:"primary_key"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
