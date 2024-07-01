package service

import (
	"github.com/juntakoman123/gin_todo_app/model"
	"github.com/juntakoman123/gin_todo_app/repository"
)

type TaskService struct {
	TaskRepo repository.TaskRepository
}

func (ts *TaskService) GetTasks() model.Tasks {
	return ts.TaskRepo.GetTasks()
}
