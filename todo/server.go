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

	err := c.BindJSON(&newTask)

	if err != nil {
		return
	}

	newTask, err = h.service.AddTask(newTask)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, newTask)

}
