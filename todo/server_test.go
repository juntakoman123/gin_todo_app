package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetTasks(t *testing.T) {

	exampleTasks := Tasks{
		{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
		{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
	}

	tests := []struct {
		name      string
		setupMock func(m *MockService)
		wantCode  int
		wantBody  string
	}{
		{
			name: "can get tasks as JSON",
			setupMock: func(m *MockService) {
				m.EXPECT().GetTasks().Return(exampleTasks, nil).Times(1)
			},
			wantCode: http.StatusOK,
			wantBody: tasksToJson(exampleTasks),
		},
		{
			name: "returns a 500 internal server error if the service fails",
			setupMock: func(m *MockService) {
				m.EXPECT().GetTasks().Return(nil, errors.New("couldn't get tasks")).Times(1)
			},
			wantCode: http.StatusInternalServerError,
			wantBody: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewMockService(ctrl)
			tt.setupMock(service)

			server := NewServer(service)

			res := httptest.NewRecorder()
			req := newGetTasksRequest()

			server.ServeHTTP(res, req)

			assertStatus(t, res, tt.wantCode)
			assertResponseBody(t, res.Body.String(), tt.wantBody)

		})
	}

}

func tasksToJson(tasks Tasks) string {
	tasksJson, _ := json.Marshal(tasks)
	return string(tasksJson)
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

func TestDeleteTask(t *testing.T) {

	t.Run("can delete task", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := 1

		service := NewMockService(ctrl)
		service.EXPECT().DeleteTask(TaskID(id)).Return(nil).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newDeleteTaskRequest(fmt.Sprint(id))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 200)
	})

	t.Run("returns 400 bad request if id param is not valid ", func(t *testing.T) {

		id := "trouble"

		server := NewServer(nil)

		res := httptest.NewRecorder()
		req := newDeleteTaskRequest(id)

		server.ServeHTTP(res, req)

		assertStatus(t, res, 400)
	})

	t.Run("returns 404 not found if task does not exist ", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := 1

		service := NewMockService(ctrl)
		service.EXPECT().DeleteTask(TaskID(id)).Return(ErrTaskNotFound).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newDeleteTaskRequest(fmt.Sprint(id))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 404)
	})

	t.Run("returns a 500 internal server error if the service fails", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := 1

		service := NewMockService(ctrl)
		service.EXPECT().DeleteTask(TaskID(id)).Return(errors.New("couldn't delete task")).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newDeleteTaskRequest(fmt.Sprint(id))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 500)
	})

}

func TestUpdateTask(t *testing.T) {

	t.Run("can update task", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := 1
		exampleTask := Task{Title: "Task 1", Status: TaskStatusTodo}
		taskJson, _ := json.Marshal(exampleTask)
		exampleTask.ID = TaskID(id)

		service := NewMockService(ctrl)
		service.EXPECT().UpdateTask(exampleTask).Return(nil).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newUpdateTaskRequest(fmt.Sprint(id), strings.NewReader(string(taskJson)))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 200)

	})

	t.Run("returns 400 bad request if id param is not valid ", func(t *testing.T) {

		exampleTask := Task{Title: "Task 1", Status: TaskStatusTodo}
		taskJson, _ := json.Marshal(exampleTask)

		server := NewServer(nil)

		res := httptest.NewRecorder()
		req := newUpdateTaskRequest("trouble", strings.NewReader(string(taskJson)))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 400)
	})

	t.Run("returns 400 bad request if body is not valid task", func(t *testing.T) {

		server := NewServer(nil)

		res := httptest.NewRecorder()
		req := newUpdateTaskRequest("1", strings.NewReader("trouble"))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 400)
	})

	t.Run("returns 404 not found if task does not exist ", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		exampleTask := Task{Title: "Task 1", Status: TaskStatusTodo}
		taskJson, _ := json.Marshal(exampleTask)

		id := 1
		exampleTask.ID = TaskID(id)

		service := NewMockService(ctrl)
		service.EXPECT().UpdateTask(exampleTask).Return(ErrTaskNotFound).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newUpdateTaskRequest(fmt.Sprint(id), strings.NewReader(string(taskJson)))

		server.ServeHTTP(res, req)

		assertStatus(t, res, 404)
	})

	t.Run("returns a 500 internal server error if the service fails", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		exampleTask := Task{Title: "Task 1", Status: TaskStatusTodo}
		taskJson, _ := json.Marshal(exampleTask)

		id := 1
		exampleTask.ID = TaskID(id)

		service := NewMockService(ctrl)
		service.EXPECT().UpdateTask(exampleTask).Return(errors.New("couldn't update task")).Times(1)

		server := NewServer(service)

		res := httptest.NewRecorder()
		req := newUpdateTaskRequest(fmt.Sprint(id), strings.NewReader(string(taskJson)))

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

func newDeleteTaskRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/tasks/%s", id), nil)
	return req
}

func newUpdateTaskRequest(id string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/tasks/%s", id), body)
	return req
}
