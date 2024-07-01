package repository

import "github.com/juntakoman123/gin_todo_app/model"

type TaskRepository interface {
	GetTasks() model.Tasks
}

type InMemoryTaskRepository struct {
	Tasks model.Tasks
}

func (repo *InMemoryTaskRepository) GetTasks() model.Tasks {
	return repo.Tasks
}
