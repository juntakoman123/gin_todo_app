package todo

type Service interface {
	GetTasks() (Tasks, error)
	AddTask(task Task) (Task, error)
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
