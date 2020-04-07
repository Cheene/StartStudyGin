package models

import (
	"StartStudyGin/22bubble/dao"
)

//表以及对应的增删改查的操作
// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Tittle string `json:"title"`
	Status bool   `json:"status"`
}

//todo 的增删改查
// CreateATodo
func CreateATodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

func GetTodoList() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodoBiId(id int) (todo *Todo, err error) {
	if err = dao.DB.Where("id=？", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}
