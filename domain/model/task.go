package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type TaskID string

func NewTaskID() TaskID {
	return TaskID(uuid.New().String())
}

type TaskTitle string

var ErrTaskTitleTooLong = errors.New("task title cannot exceed 50 characters")

func NewTaskTitle(title string) (TaskTitle, error) {
	if len(title) > 50 {
		return "", ErrTaskTitleTooLong
	}
	return TaskTitle(title), nil
}

type Task struct {
	ID      TaskID
	Title   TaskTitle
	Status  TaskStatus
	Created time.Time
	Updated time.Time
}

func NewTask(title TaskTitle) *Task {
	return &Task{
		ID:      NewTaskID(),
		Title:   title,
		Status:  TaskStatusTodo,
		Created: time.Now(),
		Updated: time.Now(),
	}
}
