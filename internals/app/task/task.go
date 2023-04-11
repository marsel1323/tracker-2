package task

import "time"

type Task struct {
	ID       int
	Name     string
	TimeGoal time.Duration
	GoalID   *int
}
