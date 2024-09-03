package usecase

import (
	"fmt"

	"github.com/juntakoman123/gin_todo_app/domain/model"
	"github.com/juntakoman123/gin_todo_app/domain/store"
)

type CreateTaskUseCase struct {
	taskStore store.Task
}

func NewCreateTaskUseCase(taskStore store.Task) *CreateTaskUseCase {
	return &CreateTaskUseCase{taskStore}
}

func (c *CreateTaskUseCase) Exec(title string) error {

	taskTitle, err := model.NewTaskTitle(title)
	if err != nil {
		return fmt.Errorf("failed to generate task title: %w", err)
	}

	task := model.NewTask(taskTitle)

	if err := c.taskStore.Insert(task); err != nil {
		return fmt.Errorf("failed to insert task to store: %w", err)
	}

	return nil
}
