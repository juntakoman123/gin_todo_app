package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juntakoman123/gin_todo_app/service"
)

type TaskController struct {
	TaskService service.TaskService
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks := tc.TaskService.GetTasks()
	c.JSON(http.StatusOK, tasks)
}
