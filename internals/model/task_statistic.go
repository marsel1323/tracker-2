package model

import "time"

type TaskStatistic struct {
	ID        int
	TaskID    int
	TimeSpent time.Duration
	CreatedAt time.Time
}
