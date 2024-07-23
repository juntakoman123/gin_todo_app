package main

import (
	"fmt"
	"os"

	"github.com/juntakoman123/gin_todo_app/config"
	"github.com/juntakoman123/gin_todo_app/todo"
)

func main() {

	store := todo.InMemoryStore{}
	service := todo.NewImplService(&store)
	server := todo.NewServer(service)

	cfg, err := config.New()
	if err != nil {
		os.Exit(1)
	}

	server.Run(fmt.Sprintf(":%d", cfg.Port))

}
