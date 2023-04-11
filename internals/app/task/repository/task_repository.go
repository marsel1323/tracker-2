package repository

import "tracker-2/internals/app/task"

type TaskRepository interface {
	CreateTask(task *task.Task) error
	FetchTaskByID(id int) (*task.Task, error)
	FetchTasks() ([]task.Task, error)
}
