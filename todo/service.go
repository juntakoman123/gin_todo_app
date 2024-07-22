package todo

type Service struct {
	store Store
}

func (s *Service) GetTasks() Tasks {
	return s.store.GetTasks()
}
