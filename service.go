package main

type Service interface {
	GetTasks() (Tasks, error)
	AddTask(task Task) (Task, error)
	DeleteTask(id TaskID) error
	UpdateTask(task Task) error
}

type ImplService struct {
	Store Store
}

func NewImplService(store Store) *ImplService {
	return &ImplService{store}
}

func (s *ImplService) GetTasks() (Tasks, error) {

	tasks, err := s.Store.GetTasks()

	if err != nil {
		return Tasks{}, err
	}

	return tasks, nil
}

func (s *ImplService) AddTask(task Task) (Task, error) {

	return Task{}, nil
}

func (s *ImplService) DeleteTask(id TaskID) error {

	return nil
}
