package repository

import (
	"time"
	"tracker-2/internals/app/task_statistic"
)

type TaskRepository interface {
	CreateTaskStatistic(task *task_statistic.TaskStatistic) error
	FetchAllTaskStatisticsByID(id int) ([]task_statistic.TaskStatistic, error)
	FetchAllTaskStatisticsForDate(taskID int, date time.Time) ([]task_statistic.TaskStatistic, error)
	FetchAllTaskStatisticsForDays(taskID int, from, to time.Time) ([]task_statistic.TaskStatistic, error)
}
