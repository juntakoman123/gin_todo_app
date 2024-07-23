package todo

type Service struct {
	store Store
}

func (s *Service) GetTasks() (Tasks, error) {
	tasks, err := s.store.GetTasks()

	if err != nil {
		return Tasks{}, err
	}

	return tasks, nil
}
