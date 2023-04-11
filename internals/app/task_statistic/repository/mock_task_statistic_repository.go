package repository

import (
	"time"
	"tracker-2/internals/app/task_statistic"
)

type mockTaskStatisticRepository struct {
	taskStatistics []task_statistic.TaskStatistic
}

func NewMockTaskStatisticRepository() *mockTaskStatisticRepository {
	return &mockTaskStatisticRepository{
		taskStatistics: []task_statistic.TaskStatistic{},
	}
}

func (m *mockTaskStatisticRepository) FetchAllTaskStatisticsByID(id int) ([]task_statistic.TaskStatistic, error) {
	taskStatistics := []task_statistic.TaskStatistic{}
	for i, statistic := range m.taskStatistics {
		if statistic.TaskID == id {
			taskStatistics = append(taskStatistics, m.taskStatistics[i])
		}
	}

	return taskStatistics, nil
}

func (m *mockTaskStatisticRepository) CreateTaskStatistic(task *task_statistic.TaskStatistic) error {
	m.taskStatistics = append(m.taskStatistics, *task)
	return nil
}

func (m *mockTaskStatisticRepository) FetchAllTaskStatisticsForDate(taskID int, date time.Time) ([]task_statistic.TaskStatistic, error) {
	taskStatistics := []task_statistic.TaskStatistic{}

	for i, statistic := range m.taskStatistics {
		if statistic.TaskID == taskID &&
			statistic.CreatedAt.Year() == date.Year() &&
			statistic.CreatedAt.Month() == date.Month() &&
			statistic.CreatedAt.Day() == date.Day() {
			taskStatistics = append(taskStatistics, m.taskStatistics[i])
		}
	}

	return taskStatistics, nil
}

func (m *mockTaskStatisticRepository) FetchAllTaskStatisticsForDays(taskID int, from, to time.Time) ([]task_statistic.TaskStatistic, error) {
	taskStatistics := []task_statistic.TaskStatistic{}

	for i, statistic := range m.taskStatistics {
		if statistic.TaskID == taskID &&
			statistic.CreatedAt.After(from) &&
			statistic.CreatedAt.Before(to) {
			taskStatistics = append(taskStatistics, m.taskStatistics[i])
		}
	}

	return taskStatistics, nil
}
