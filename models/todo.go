package models

import "Weblist-go-backend/dao"

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
	Todo 增删改查
*/

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodolist() (todolist []*Todo, err error) {
	if err = dao.DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("id = ?", id).Find(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodoById(id string, todo *Todo) (err error) {
	err = dao.DB.Model(&Todo{}).Where("id = ?", id).Update(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Delete(&Todo{}, id).Error
	return
}
