package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
		setupMock func(m *MockTasKService)
		wantCode  int
		wantBody  string
	}{
		{
			name: "can get tasks as JSON",
			setupMock: func(m *MockTasKService) {
				m.EXPECT().GetTasks().Return(exampleTasks, nil).Times(1)
			},
			wantCode: http.StatusOK,
			wantBody: toJSON(exampleTasks),
		},
		{
			name: "returns a 500 internal server error if the service fails",
			setupMock: func(m *MockTasKService) {
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

			service := NewMockTasKService(ctrl)
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

func TestPostTask(t *testing.T) {

	inputTask := Task{Title: "Task 1", Status: TaskStatusTodo}
	wantTask := Task{ID: 1, Title: "Task 1", Status: TaskStatusTodo}

	tests := []struct {
		name      string
		reqBody   string
		setupMock func(m *MockTasKService)
		wantCode  int
		wantBody  string
	}{
		{
			name:    "can add valid task",
			reqBody: toJSON(inputTask),
			setupMock: func(m *MockTasKService) {
				m.EXPECT().AddTask(inputTask).Return(wantTask, nil).Times(1)
			},
			wantCode: http.StatusOK,
			wantBody: toJSON(wantTask),
		},
		{
			name:    "returns 400 bad request if body is not valid task JSON",
			reqBody: "trouble",
			setupMock: func(m *MockTasKService) {
				m.EXPECT().AddTask(gomock.Any()).Times(0)
			},
			wantCode: http.StatusBadRequest,
			wantBody: "",
		},
		{
			name:    "returns a 500 internal server error if the service fails",
			reqBody: toJSON(inputTask),
			setupMock: func(m *MockTasKService) {
				m.EXPECT().AddTask(inputTask).Return(Task{}, errors.New("couldn't add new task")).Times(1)
			},
			wantCode: http.StatusInternalServerError,
			wantBody: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewMockTasKService(ctrl)
			tt.setupMock(service)

			server := NewServer(service)

			res := httptest.NewRecorder()
			req := newPostTaskRequest(tt.reqBody)

			server.ServeHTTP(res, req)

			assertStatus(t, res, tt.wantCode)
			assertResponseBody(t, res.Body.String(), tt.wantBody)

		})
	}

}

func TestDeleteTask(t *testing.T) {

	tests := []struct {
		name      string
		id        string
		setupMock func(m *MockTasKService)
		wantCode  int
	}{
		{
			name: "can delete task",
			id:   "1",
			setupMock: func(m *MockTasKService) {
				m.EXPECT().DeleteTask(TaskID(1)).Return(nil).Times(1)
			},
			wantCode: http.StatusOK,
		},
		{
			name: "returns 400 bad request if id param is not valid",
			id:   "trouble",
			setupMock: func(m *MockTasKService) {
				m.EXPECT().AddTask(gomock.Any()).Times(0)
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "returns 404 not found if task does not exist",
			id:   "1",
			setupMock: func(m *MockTasKService) {
				m.EXPECT().DeleteTask(TaskID(1)).Return(ErrTaskNotFound).Times(1)
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "returns a 500 internal server error if the service fails",
			id:   "1",
			setupMock: func(m *MockTasKService) {
				m.EXPECT().DeleteTask(TaskID(1)).Return(errors.New("couldn't delete task")).Times(1)
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewMockTasKService(ctrl)
			tt.setupMock(service)

			server := NewServer(service)

			res := httptest.NewRecorder()
			req := newDeleteTaskRequest(tt.id)

			server.ServeHTTP(res, req)

			assertStatus(t, res, tt.wantCode)
		})
	}

}

func TestUpdateTask(t *testing.T) {

	inputTask := Task{Title: "Task 1", Status: TaskStatusTodo}

	tests := []struct {
		name      string
		id        string
		reqBody   string
		setupMock func(m *MockTasKService)
		wantCode  int
	}{
		{
			name:    "can update task",
			id:      "1",
			reqBody: toJSON(inputTask),
			setupMock: func(m *MockTasKService) {
				m.EXPECT().UpdateTask(Task{ID: 1, Title: "Task 1", Status: TaskStatusTodo}).Return(nil).Times(1)
			},
			wantCode: http.StatusOK,
		},
		{
			name:    "returns 400 bad request if id param is not valid",
			id:      "trouble",
			reqBody: toJSON(inputTask),
			setupMock: func(m *MockTasKService) {
				m.EXPECT().UpdateTask(gomock.Any()).Times(0)
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name:    "returns 400 bad request if body is not valid task",
			id:      "1",
			reqBody: "trouble",
			setupMock: func(m *MockTasKService) {
				m.EXPECT().UpdateTask(gomock.Any()).Times(0)
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name:    "returns 404 not found if task does not exist",
			id:      "1",
			reqBody: toJSON(inputTask),
			setupMock: func(m *MockTasKService) {
				m.EXPECT().UpdateTask(Task{ID: 1, Title: "Task 1", Status: TaskStatusTodo}).Return(ErrTaskNotFound).Times(1)
			},
			wantCode: http.StatusNotFound,
		},
		{
			name:    "returns a 500 internal server error if the service fails",
			id:      "1",
			reqBody: toJSON(inputTask),
			setupMock: func(m *MockTasKService) {
				m.EXPECT().UpdateTask(Task{ID: 1, Title: "Task 1", Status: TaskStatusTodo}).Return(errors.New("couldn't update task")).Times(1)
			},
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewMockTasKService(ctrl)
			tt.setupMock(service)

			server := NewServer(service)

			res := httptest.NewRecorder()
			req := newUpdateTaskRequest(tt.id, tt.reqBody)

			server.ServeHTTP(res, req)

			assertStatus(t, res, tt.wantCode)
		})
	}

}

func toJSON(v any) string {
	json, _ := json.Marshal(v)
	return string(json)
}

func newGetTasksRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	return req
}

func newPostTaskRequest(body string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/tasks", strings.NewReader(body))
	return req
}

func newDeleteTaskRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/tasks/%s", id), nil)
	return req
}

func newUpdateTaskRequest(id string, body string) *http.Request {
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/tasks/%s", id), strings.NewReader(body))
	return req
}
