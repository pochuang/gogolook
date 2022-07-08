package utest

import (
	"internal/entity"
	"internal/model"
)

func setupTaskData() (data []model.Task) {
	task1 := model.Task{Name: "Buy breakfast"}
	task2 := model.Task{Name: "Buy lunch"}
	task3 := model.Task{Name: "Buy dinner"}
	data = append(data, task1)
	data = append(data, task2)
	data = append(data, task3)
	return
}

func tearDown(data []model.Task) {
	db := entity.OpenDatabase()
	db.Clear(data)
}
