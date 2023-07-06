package utest

import (
	"internal/entity"
	"testing"
)

func TestCreate(t *testing.T) {
	db := entity.OpenDatabase()
	data := setupTaskData()
	defer func() {
		tearDown(data)
	}()
	db.Create(data...)
	tasks := db.List()
	for i := range tasks {
		if !tasks[i].CheckValue(data[i]) {
			t.Error("data not created")
		}
	}
}

func TestDelete(t *testing.T) {
	db := entity.OpenDatabase()
	data := setupTaskData()
	defer func() {
		tearDown(data)
	}()
	created := db.Create(data...)
	db.Del(created[0].ID)
	tasks := db.List()
	for _, task := range tasks {
		if task.Name == data[0].Name {
			t.Error("data not deleted")
		}
	}
}
