package service

import (
	"internal/entity"
	"internal/model"
)

// Task is a service that can do CRUD for task.
type Task struct {
	Name   string
	Status bool
}

// List will return all data.
func (t *Task) List() (tasks []model.Task) {
	db := entity.OpenDatabase()
	return db.List()
}

// Create can be used to append a row to database.
func (t *Task) Create(task model.Task) []model.Task {
	db := entity.OpenDatabase()
	return db.Create(task)
}

// Update is used to update data for specified row in database.
func (t *Task) Update(id int64, task model.Task) (updatedTask model.Task, ok bool) {
	db := entity.OpenDatabase()
	return db.Update(id, task)
}

// Delete will remove a row from database.
func (t *Task) Delete(id int64) {
	db := entity.OpenDatabase()
	db.Del(id)
	return
}
