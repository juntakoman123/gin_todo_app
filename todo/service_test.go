package todo

import (
	"reflect"
	"testing"
)

type stubStore struct {
	tasks Tasks
}

func (s *stubStore) GetTasks() Tasks {
	return s.tasks
}

func TestService(t *testing.T) {

	want := Tasks{
		{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
		{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
	}

	store := stubStore{want}

	service := Service{&store}
	got := service.GetTasks()

	assertTasks(t, got, want)
}

func assertTasks(t testing.TB, got, want Tasks) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
