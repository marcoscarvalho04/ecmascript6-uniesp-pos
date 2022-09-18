package list

import "gorm.io/gorm"

type TodoListModel struct {
	model  gorm.Model
	Id     int32  `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (TodoListModel) TableName() string {
	return "list"
}
