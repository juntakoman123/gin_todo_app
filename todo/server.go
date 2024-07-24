package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(service Service) *gin.Engine {

	handler := handler{service}

	r := gin.Default()
	r.GET("/tasks", handler.GetTasks)
	r.POST("/tasks", handler.PostTask)

	return r

}

type handler struct {
	service Service
}

func (h *handler) GetTasks(c *gin.Context) {

	tasks, err := h.service.GetTasks()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tasks)

}

func (h *handler) PostTask(c *gin.Context) {

	var newTask Task

	c.BindJSON(&newTask)

	newTask, _ = h.service.AddTask(newTask)

	c.JSON(http.StatusOK, newTask)

}
