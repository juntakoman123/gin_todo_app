package main

import (
	"fmt"
	"os"

	"github.com/juntakoman123/gin_todo_app/config"
	"github.com/juntakoman123/gin_todo_app/todo"
)

var tasks = todo.Tasks{
	{ID: 1, Title: "Task 1", Status: todo.TaskStatusTodo},
	{ID: 2, Title: "Task 2", Status: todo.TaskStatusTodo},
}

func main() {

	store := todo.NewinMemoryStore(tasks)
	service := todo.NewImplService(store)
	server := todo.NewServer(service)

	cfg, err := config.New()
	if err != nil {
		os.Exit(1)
	}

	server.Run(fmt.Sprintf(":%d", cfg.Port))

}
