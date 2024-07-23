package todo

type Store interface {
	GetTasks() (Tasks, error)
}

type inMemoryStore struct {
	tasks Tasks
}

func NewinMemoryStore(tasks Tasks) *inMemoryStore {
	return &inMemoryStore{tasks}
}

func (s *inMemoryStore) GetTasks() (Tasks, error) {
	return s.tasks, nil
}
