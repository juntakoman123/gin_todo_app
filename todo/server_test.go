package todo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type mockService struct {
	getTasksFunc func() (Tasks, error)
	postTaskFunc func(task Task) (Task, error)
	tasksAdded   Tasks
}

func (s *mockService) GetTasks() (Tasks, error) {
	return s.getTasksFunc()
}

func (s *mockService) AddTask(task Task) (Task, error) {

	s.tasksAdded = append(s.tasksAdded, task)
	return s.postTaskFunc(task)
}

func TestGetTasks(t *testing.T) {

	t.Run("can get tasks as JSON", func(t *testing.T) {

		exampleTasks := Tasks{
			{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
			{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
		}

		tasksJson, _ := json.Marshal(exampleTasks)

		service := mockService{
			getTasksFunc: func() (Tasks, error) {
				return exampleTasks, nil
			},
		}

		server := NewServer(&service)

		res := httptest.NewRecorder()
		req := newGetTasksRequest()

		server.ServeHTTP(res, req)

		assertStatus(t, res, 200)

		assertResponseBody(t, res.Body.String(), string(tasksJson))

	})

	t.Run("returns a 500 internal server error if the service fails", func(t *testing.T) {

		service := mockService{
			getTasksFunc: func() (Tasks, error) {
				return Tasks{}, errors.New("couldn't get tasks")
			},
		}

		server := NewServer(&service)

		res := httptest.NewRecorder()
		req := newGetTasksRequest()

		server.ServeHTTP(res, req)

		assertStatus(t, res, 500)

	})

}

func TestPostTask(t *testing.T) {

	t.Run("can add valid task", func(t *testing.T) {
		exampleTask := Task{Title: "Task 1", Status: TaskStatusTodo}
		taskJson, _ := json.Marshal(exampleTask)

		exampleInsertedID := 1
		wantTask := Task{ID: TaskID(exampleInsertedID), Title: "Task 1", Status: TaskStatusTodo}
		wantTaskJson, _ := json.Marshal(wantTask)

		service := mockService{
			postTaskFunc: func(task Task) (Task, error) {
				task.ID = TaskID(exampleInsertedID)
				return task, nil
			},
		}

		server := NewServer(&service)

		res := httptest.NewRecorder()
		req := newPostTaskRequest(strings.NewReader(string(taskJson)))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 200)

		if len(service.tasksAdded) != 1 {
			t.Fatalf("expected 1 task added but got %d", len(service.tasksAdded))
		}

		if !reflect.DeepEqual(service.tasksAdded[0], exampleTask) {
			t.Errorf("expected %v posted 1 but %v posted", exampleTask, service.tasksAdded[0])
		}

		assertResponseBody(t, res.Body.String(), string(wantTaskJson))

	})
}

func newGetTasksRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	return req
}

func newPostTaskRequest(body io.Reader) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/tasks", body)
	return req
}
