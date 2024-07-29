package main

type Store interface {
	GetTasks() (Tasks, error)
}
