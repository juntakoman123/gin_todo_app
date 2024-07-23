package todo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubService struct {
	tasks Tasks
}

func (s *stubService) GetTasks() (Tasks, error) {
	return s.tasks, nil
}

func TestGetTasks(t *testing.T) {

	t.Run("it returns tasks as JSON", func(t *testing.T) {

		exampleTasks := Tasks{
			{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
			{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
		}

		tasksJson, _ := json.Marshal(exampleTasks)

		service := stubService{exampleTasks}
		server := NewServer(&service)

		res := httptest.NewRecorder()
		req := newGetTasksRequest()

		server.ServeHTTP(res, req)

		assertStatus(t, res, 200)

		assertResponseBody(t, res.Body.String(), string(tasksJson))

	})

}

func newGetTasksRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	return req
}
