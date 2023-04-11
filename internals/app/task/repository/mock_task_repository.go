package repository

import "tracker-2/internals/app/task"

type mockTaskRepository struct {
	tasks []task.Task
}

func NewMockTaskRepository() *mockTaskRepository {
	return &mockTaskRepository{
		tasks: []task.Task{},
	}
}

func (m *mockTaskRepository) FetchTasks() ([]task.Task, error) {
	return m.tasks, nil
}

func (m *mockTaskRepository) CreateTask(task *task.Task) error {
	m.tasks = append(m.tasks, *task)
	return nil
}

func (m *mockTaskRepository) FetchTaskByID(id int) (*task.Task, error) {
	for _, t := range m.tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, nil
}
