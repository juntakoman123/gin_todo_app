package todo

type Store interface {
	GetTasks() (Tasks, error)
}
