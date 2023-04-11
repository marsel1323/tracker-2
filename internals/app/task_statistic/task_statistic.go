package task_statistic

import "time"

type TaskStatistic struct {
	ID        int
	TimeSpent time.Duration
	TaskID    int
	CreatedAt time.Time
}
