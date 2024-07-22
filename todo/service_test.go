package todo

import (
	"errors"
	"reflect"
	"testing"
)

type stubStore struct {
	tasks Tasks
}

func (s *stubStore) GetTasks() (Tasks, error) {
	return s.tasks, nil
}

type errStubStore struct {
	err error
}

func (e *errStubStore) GetTasks() (Tasks, error) {
	return Tasks{}, e.err
}

func TestService(t *testing.T) {

	t.Run("get tasks", func(t *testing.T) {
		want := Tasks{
			{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
			{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
		}

		store := stubStore{want}

		service := Service{&store}
		got, err := service.GetTasks()

		assertNoError(t, err)
		assertTasks(t, got, want)
	})

	t.Run("it returns db error", func(t *testing.T) {

		want := Tasks{}
		wantErr := errors.New("db error")

		store := errStubStore{wantErr}
		service := Service{&store}

		got, err := service.GetTasks()

		assertError(t, err, wantErr)
		assertTasks(t, got, want)

	})

}

func assertTasks(t testing.TB, got, want Tasks) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
