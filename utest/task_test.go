package utest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gogolook/api"
	"internal/model"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	task := api.TaskHandler{}
	router.GET("/task", task.List)
	router.POST("/task", task.Create)
	router.DELETE("/task/:id", task.Delete)
	router.PUT("/task/:id", task.Update)
	return router
}

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func convertMapToTask(data map[string]interface{}) (result model.Task, err error) {
	jsonBody, err1 := json.Marshal(data)
	if err1 != nil {
		err = err1
		return
	}
	err = json.Unmarshal(jsonBody, &result)
	return
}

func createData(t *testing.T, tasks []model.Task, router *gin.Engine) (createdTasks []model.Task) {
	// Create data for test
	for _, task := range tasks {
		req, _ := json.Marshal(task)
		w := performRequest(router, "POST", "/task", req)
		statusCode := w.Code
		if statusCode != http.StatusCreated {
			t.Error("failed to create task")
		}
		response := model.Response{}
		json.NewDecoder(w.Body).Decode(&response)
		result := response.Result.(map[string]interface{})
		if created, err := convertMapToTask(result); err == nil {
			createdTasks = append(createdTasks, created)
		}
	}
	return
}

func TestCreateTask(t *testing.T) {
	tasks := setupTaskData()
	router := setupRouter()
	defer func() {
		tearDown(tasks)
	}()
	// We will run this test parallel
	t.Parallel()
	loop := 100
	for i := 0; i < loop; i++ {
		t.Run("Loop"+strconv.Itoa(i), func(t1 *testing.T) {
			for _, task := range tasks {
				// 1. Create data
				task.Name = fmt.Sprintf("%s-%d", task.Name, i)
				req, _ := json.Marshal(task)
				w := performRequest(router, "POST", "/task", req)
				// 2. Get data from response
				response := model.Response{}
				json.NewDecoder(w.Body).Decode(&response)
				result := response.Result.(map[string]interface{})
				// 3. Compare if they are equal
				if created, err := convertMapToTask(result); err != nil || !created.CheckValue(task) {
					t.Error("failed to create")
				}
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := setupTaskData()
	router := setupRouter()
	defer func() {
		tearDown(tasks)
	}()

	// 1. Create data for test
	createdTasks := createData(t, tasks, router)

	// 2. Delete first row
	first := createdTasks[0]
	performRequest(router, "DELETE", fmt.Sprintf("/task/%d", first.ID), nil)

	// 3. Get all data and check if data is deleted
	w := performRequest(router, "GET", "/task", nil)
	response := model.Response{}
	json.NewDecoder(w.Body).Decode(&response)
	result := response.Result.([]interface{})
	for _, r := range result {
		data1 := r.(map[string]interface{})

		if task, err := convertMapToTask(data1); err != nil {
			t.Error("data format error")
		} else {
			if task.CheckValue(tasks[0]) {
				t.Error("data not deleted")
			}
		}
	}
}

func TestUpdateTask(t *testing.T) {
	tasks := setupTaskData()
	router := setupRouter()
	defer func() {
		tearDown(tasks)
	}()
	// 1. Create data for test
	createdTasks := createData(t, tasks, router)

	// 2. Do some modifications to data
	for i := range createdTasks {
		createdTasks[i].Status = 1
		req, _ := json.Marshal(createdTasks[i])
		performRequest(router, "PUT", fmt.Sprintf("/task/%d", createdTasks[i].ID), req)
	}

	// 3. Check data if their status = 1
	w := performRequest(router, "GET", "/task", nil)
	response := model.Response{}
	json.NewDecoder(w.Body).Decode(&response)
	result := response.Result.([]interface{})
	for _, r := range result {
		data1 := r.(map[string]interface{})

		if task, _ := convertMapToTask(data1); task.Status != 1 {
			t.Error("status not updated")
		}
	}
}
