package todo

type Service interface {
	GetTasks() (Tasks, error)
}

type ImplService struct {
	store Store
}

func (s *ImplService) GetTasks() (Tasks, error) {
	tasks, err := s.store.GetTasks()

	if err != nil {
		return Tasks{}, err
	}

	return tasks, nil
}
