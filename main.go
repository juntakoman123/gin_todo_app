package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/juntakoman123/gin_todo_app/config"
)

func main() {

	router := gin.Default()
	// router.GET("/tasks", todo.handlerGetTasks)

	cfg, err := config.New()
	if err != nil {
		os.Exit(1)
	}

	router.Run(fmt.Sprintf(":%d", cfg.Port))

}
