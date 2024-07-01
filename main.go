package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/juntakoman123/gin_todo_app/config"
	"github.com/juntakoman123/gin_todo_app/controller"
	"github.com/juntakoman123/gin_todo_app/model"
	"github.com/juntakoman123/gin_todo_app/repository"
	"github.com/juntakoman123/gin_todo_app/service"
)

func main() {

	taskRepo := &repository.InMemoryTaskRepository{
		Tasks: model.Tasks{
			{ID: 1, Title: "Task 1", Status: model.TaskStatusTodo},
			{ID: 2, Title: "Task 2", Status: model.TaskStatusTodo},
		},
	}

	taskService := service.TaskService{
		TaskRepo: taskRepo,
	}

	taskController := &controller.TaskController{
		TaskService: taskService,
	}

	router := gin.Default()

	router.GET("/tasks", taskController.GetTasks)

	cfg, err := config.New()
	if err != nil {
		os.Exit(1)
	}

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
