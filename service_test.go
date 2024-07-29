package main

import (
	"errors"
	"testing"
)

type stubStore struct {
	tasks Tasks
}

func (s *stubStore) GetTasks() (Tasks, error) {
	return s.tasks, nil
}

type errStubStore struct {
	tasks Tasks
	err   error
}

func (e *errStubStore) GetTasks() (Tasks, error) {
	return e.tasks, e.err
}

func TestService(t *testing.T) {

	t.Run("get tasks", func(t *testing.T) {
		want := Tasks{
			{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
			{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
		}

		store := stubStore{want}

		service := ImplService{&store}
		got, err := service.GetTasks()

		assertNoError(t, err)
		assertTasks(t, got, want)
	})

	t.Run("it returns db error", func(t *testing.T) {

		want := Tasks{}
		wantErr := errors.New("db error")

		store := errStubStore{Tasks{
			{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
			{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
		}, wantErr}

		service := ImplService{&store}

		got, err := service.GetTasks()

		assertError(t, err, wantErr)
		assertTasks(t, got, want)

	})

}
