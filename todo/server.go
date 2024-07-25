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

	task, err := h.parseTask(c)

	if err != nil {
		return
	}

	newTask, err := h.service.AddTask(task)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, newTask)

}

func (h *handler) DeleteTask(c *gin.Context) {

	id, err := h.parseID(c)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(id)

	if errors.Is(err, ErrTaskNotFound) {
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

	id, err := h.parseID(c)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	updateTask, err := h.parseTask(c)

	if err != nil {
		return
	}

	updateTask.ID = id

	err = h.service.UpdateTask(updateTask)

	if errors.Is(err, ErrTaskNotFound) {
		c.Status(http.StatusNotFound)
		return
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *handler) parseID(c *gin.Context) (TaskID, error) {
	rawId := c.Param("id")
	id, err := strconv.Atoi(rawId)

	if err != nil {
		return 0, err
	}

	return TaskID(id), nil
}

func (h *handler) parseTask(c *gin.Context) (Task, error) {
	var task Task

	if err := c.BindJSON(&task); err != nil {
		return Task{}, err
	}

	return task, nil
}
