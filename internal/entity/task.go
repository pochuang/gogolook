package entity

import (
	"internal/model"
	"sort"
	"sync"
)

var privateData sync.Map
var autoIncremental int64 = 0 // auto incremental id
var locker sync.Mutex
var dbOnce sync.Once
var dbInstance Mock

// Mock is a simple in-memory database for coding test of GogoLook.
type Mock struct {
}

func OpenDatabase() Mock {
	dbOnce.Do(func() {
		dbInstance = Mock{}
	})
	return dbInstance
}

// Create will add a row to database
func (d *Mock) Create(data ...model.Task) (result []model.Task) {
	defer locker.Unlock()
	locker.Lock()
	for _, task := range data {
		autoIncremental++
		task.ID = autoIncremental
		privateData.Store(autoIncremental, task)
		result = append(result, task)
	}
	return
}

// Del will remove a row from database
func (d *Mock) Del(id int64) {
	privateData.Delete(id)
}

// Update will update a row to database
func (d *Mock) Update(id int64, task model.Task) (updatedTask model.Task, ok bool) {
	if oldTask, ok := privateData.Load(id); ok {
		prev, _ := oldTask.(model.Task)
		if prev.ID > 0 {
			task.ID = prev.ID
			if task.Name != "" {
				prev.Name = task.Name
			} else {
				task.Name = prev.Name
			}
			privateData.Store(id, task)
			return task, true
		}
	}
	return task, false
}

// Get will obtain a row from database
func (d *Mock) Get(id int64) (result string, ok bool) {
	if data, ok1 := privateData.Load(id); ok1 {
		result, _ = data.(string)
		ok = ok1
	}

	return
}

// Clear will empty our database
func (d *Mock) Clear(tasks []model.Task) {
	for _, task := range tasks {
		privateData.Delete(task.ID)
	}
}

// List will return all data from database
func (d *Mock) List() (tasks []model.Task) {
	privateData.Range(func(key interface{}, value interface{}) bool {
		if task, ok1 := value.(model.Task); ok1 {
			tasks = append(tasks, task)
			return true
		}
		return false
	})
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})
	return
}
