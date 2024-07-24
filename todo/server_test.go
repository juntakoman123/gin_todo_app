package todo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetTasks(t *testing.T) {

	t.Run("can get tasks as JSON", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		exampleTasks := Tasks{
			{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
			{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
		}

		tasksJson, _ := json.Marshal(exampleTasks)

		service := NewMockService(ctrl)
		service.EXPECT().GetTasks().Return(exampleTasks, nil).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newGetTasksRequest()

		server.ServeHTTP(res, req)

		assertStatus(t, res, 200)

		assertResponseBody(t, res.Body.String(), string(tasksJson))

	})

	t.Run("returns a 500 internal server error if the service fails", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := NewMockService(ctrl)
		service.EXPECT().GetTasks().Return(Tasks{}, errors.New("couldn't get tasks")).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newGetTasksRequest()

		server.ServeHTTP(res, req)

		assertStatus(t, res, 500)

	})

}

func TestPostTask(t *testing.T) {

	t.Run("can add valid task", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		exampleTask := Task{Title: "Task 1", Status: TaskStatusTodo}
		taskJson, _ := json.Marshal(exampleTask)

		wantTask := Task{ID: 1, Title: "Task 1", Status: TaskStatusTodo}
		wantTaskJson, _ := json.Marshal(wantTask)

		service := NewMockService(ctrl)
		service.EXPECT().AddTask(exampleTask).Return(wantTask, nil).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newPostTaskRequest(strings.NewReader(string(taskJson)))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 200)
		assertResponseBody(t, res.Body.String(), string(wantTaskJson))

	})

	t.Run("returns 400 bad request if body is not valid task JSON", func(t *testing.T) {

		server := NewServer(nil)

		res := httptest.NewRecorder()
		req := newPostTaskRequest(strings.NewReader("trouble"))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 400)
	})

	t.Run("returns a 500 internal server error if the service fails", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		exampleTask := Task{Title: "Task 1", Status: TaskStatusTodo}
		taskJson, _ := json.Marshal(exampleTask)

		service := NewMockService(ctrl)
		service.EXPECT().AddTask(exampleTask).Return(Task{}, errors.New("couldn't add new task")).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newPostTaskRequest(strings.NewReader(string(taskJson)))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 500)
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
