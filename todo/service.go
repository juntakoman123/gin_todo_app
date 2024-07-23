package todo

type Service interface {
	GetTasks() (Tasks, error)
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
