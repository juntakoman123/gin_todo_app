package todo

type Service struct {
	store Store
}

func (s *Service) GetTasks() (Tasks, error) {
	return s.store.GetTasks()
}
