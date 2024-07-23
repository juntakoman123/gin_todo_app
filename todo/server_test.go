package todo

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
	getTasksFunc func() (Tasks, error)
}

func (s *mockService) GetTasks() (Tasks, error) {
	return s.getTasksFunc()
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

func newGetTasksRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	return req
}
