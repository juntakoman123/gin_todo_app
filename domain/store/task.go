package store

import "github.com/juntakoman123/gin_todo_app/domain/model"

type Task interface {
	Insert(task *model.Task) error
}
