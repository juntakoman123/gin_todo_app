package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(service Service) *gin.Engine {

	handler := handler{service}

	r := gin.Default()
	r.GET("/tasks", handler.GetTasks)

	return r

}

type handler struct {
	service Service
}

func (h *handler) GetTasks(c *gin.Context) {

	tasks, _ := h.service.GetTasks()

	c.JSON(http.StatusOK, tasks)

}
