package main

import (
	"github.com/juntakoman123/gin_todo_app/todo"
)

var tasks = todo.Tasks{
	{ID: 1, Title: "Task 1", Status: todo.TaskStatusTodo},
	{ID: 2, Title: "Task 2", Status: todo.TaskStatusTodo},
}

type inMemoryStore struct {
	tasks todo.Tasks
}

func NewinMemoryStore(tasks todo.Tasks) *inMemoryStore {
	return &inMemoryStore{tasks}
}

func (s *inMemoryStore) GetTasks() (todo.Tasks, error) {
	return s.tasks, nil
}

func main() {

	// store := NewinMemoryStore(tasks)
	// service := todo.NewImplService(store)
	// server := todo.NewServer(service)

	// cfg, err := config.New()
	// if err != nil {
	// 	os.Exit(1)
	// }

	// server.Run(fmt.Sprintf(":%d", cfg.Port))

}
