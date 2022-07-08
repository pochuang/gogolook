package main

//
// Author: Eric Lin
// Date: 2022/07/07
// Version: 1.0.0
// Desc: This project is for GogoLook demo purpose only.

import (
	"fmt"
	"gogolook/api"

	"github.com/gin-gonic/gin"
)

// setupRouter will get routers ready to go.
func setupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	router := gin.New()
	task := api.TaskHandler{}
	router.GET("/task", task.List)
	router.POST("/task", task.Create)
	router.DELETE("/task/:id", task.Delete)
	router.PUT("/task/:id", task.Update)
	return router
}

// Running as a service
func main() {
	router := setupRouter()
	fmt.Println("GOGOLook service is started!")
	router.Run(":8080")
}
