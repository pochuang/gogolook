package api

import (
	"internal/model"
	"internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	svc service.Task
}

// List gets all data from our database in memory.
func (t *TaskHandler) List(c *gin.Context) {
	response := model.Response{}
	tasks := t.svc.List()
	status := http.StatusOK

	if len(tasks) == 0 {
		response.Result = []model.Task{}
		status = http.StatusNotFound
	} else {
		response.Result = tasks
	}

	c.JSON(status, response)
	return
}

// Create will add a row.
func (t *TaskHandler) Create(c *gin.Context) {
	response := model.Response{}
	task := model.Task{}
	if err := c.BindJSON(&task); err != nil {
		return
	}
	created := t.svc.Create(task)
	status := http.StatusCreated
	if len(created) == 1 {
		response.Result = created[0]
	} else {
		// No data created
		response.Result = task
	}

	c.JSON(status, response)
}

//  Update will update our data exist in database.
func (t *TaskHandler) Update(c *gin.Context) {
	response := model.Response{}
	id := c.Param("id")
	task := model.Task{}
	if err := c.BindJSON(&task); err != nil {
		return
	}

	taskID, _ := strconv.ParseInt(id, 10, 64)
	updated, _ := t.svc.Update(taskID, task)

	httpStatus := http.StatusOK
	response.Result = updated
	c.JSON(httpStatus, response)
}

// Delete can be used to remove a row by it's id.
func (t *TaskHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	taskID, _ := strconv.ParseInt(id, 10, 64)
	t.svc.Delete(taskID)
	c.Status(http.StatusOK)
}
