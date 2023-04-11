package model

import "time"

type Goal struct {
	ID           int
	Title        string
	Description  string
	StartDate    time.Time
	EndDate      time.Time
	CategoryID   *int
	ParentGoalID *int
}
