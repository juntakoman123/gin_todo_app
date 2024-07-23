package todo

type Store interface {
	GetTasks() (Tasks, error)
}

type InMemoryStore struct {
}

func (s *InMemoryStore) GetTasks() (Tasks, error) {

	return Tasks{
		{ID: 1, Title: "Task 1", Status: TaskStatusTodo},
		{ID: 2, Title: "Task 2", Status: TaskStatusTodo},
	}, nil

}
