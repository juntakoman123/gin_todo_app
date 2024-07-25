package todo

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrTaskNotFound = errors.New("task not found")

func NewServer(service Service) *gin.Engine {

	handler := handler{service}

	r := gin.Default()
	r.GET("/tasks", handler.GetTasks)
	r.POST("/tasks", handler.PostTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)
	r.PUT("/tasks/:id", handler.UpdateTask)

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

func (h *handler) DeleteTask(c *gin.Context) {

	rawId := c.Param("id")
	id, err := strconv.Atoi(rawId)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(TaskID(id))

	if err == ErrTaskNotFound {
		c.Status(http.StatusNotFound)
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}

func (h *handler) UpdateTask(c *gin.Context) {

	rawId := c.Param("id")
	id, err := strconv.Atoi(rawId)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var updateTask Task
	err = c.BindJSON(&updateTask)

	if err != nil {
		return
	}

	updateTask.ID = TaskID(id)

	err = h.service.UpdateTask(updateTask)

	if err == ErrTaskNotFound {
		c.Status(http.StatusNotFound)
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
