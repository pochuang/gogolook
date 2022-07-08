package main

//
// Author: Eric Lin
// Date: 2022/07/07
// Version: 1.0.0
// Desc: This project is for GogoLook demo purpose only.

import (
	"context"
	"fmt"
	"gogolook/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// setupRouter will get routers ready to go.
func setupRouter() *gin.Engine {
	router := gin.Default()
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
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Print("failed to listen:", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown with error:", err)
	}
	<-ctx.Done()
}
