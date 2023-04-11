package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"tracker-2/internals/app/task"
	"tracker-2/internals/app/task/repository"
	"tracker-2/internals/app/task_statistic"
)

func TestMockTaskStatisticRepository_CreateTaskStatistic(t *testing.T) {
	morningWorkout, err := CreateMorningWorkoutTask(t)

	taskStatisticRepo := NewMockTaskStatisticRepository()

	CreateMorningWorkoutStatistics(t, morningWorkout, err, taskStatisticRepo)
}

func TestMockTaskStatisticRepository_FetchAllTaskStatisticsByID(t *testing.T) {
	morningWorkout, err := CreateMorningWorkoutTask(t)

	taskStatisticRepo := NewMockTaskStatisticRepository()

	morningWorkoutStatistics := CreateMorningWorkoutStatistics(t, morningWorkout, err, taskStatisticRepo)

	fetchedTaskStatistics, err := taskStatisticRepo.FetchAllTaskStatisticsByID(morningWorkout.ID)
	assert.NoError(t, err)
	assert.Equal(t, morningWorkoutStatistics, fetchedTaskStatistics)
}

// проверить сколько за определенный/конкретный день статистики вернется
func TestMockTaskStatisticRepository_FetchAllTaskStatisticsForToday(t *testing.T) {
	morningWorkout, err := CreateMorningWorkoutTask(t)

	taskStatisticRepo := NewMockTaskStatisticRepository()

	morningWorkoutStatistics := CreateMorningWorkoutStatistics(t, morningWorkout, err, taskStatisticRepo)

	date := time.Date(2023, 4, 11, 10, 48, 0, 0, time.UTC)
	fetchedTaskStatistics, err := taskStatisticRepo.FetchAllTaskStatisticsForDate(morningWorkout.ID, date)

	assert.NoError(t, err)
	assert.Equal(t, []task_statistic.TaskStatistic{morningWorkoutStatistics[0]}, fetchedTaskStatistics)

	date = time.Date(2023, 4, 10, 10, 48, 0, 0, time.UTC)
	fetchedTaskStatistics, err = taskStatisticRepo.FetchAllTaskStatisticsForDate(morningWorkout.ID, date)

	assert.NoError(t, err)
	assert.Equal(t, []task_statistic.TaskStatistic{morningWorkoutStatistics[1]}, fetchedTaskStatistics)

}

// сколько за все прошедшие дни
func TestMockTaskStatisticRepository_FetchAllTaskStatisticsForDays(t *testing.T) {
	morningWorkout, err := CreateMorningWorkoutTask(t)

	taskStatisticRepo := NewMockTaskStatisticRepository()

	morningWorkoutStatistics := CreateMorningWorkoutStatistics(t, morningWorkout, err, taskStatisticRepo)

	from := time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC)
	to := time.Date(2023, 4, 12, 0, 0, 0, 0, time.UTC)
	fetchedTaskStatistics, err := taskStatisticRepo.FetchAllTaskStatisticsForDays(morningWorkout.ID, from, to)
	assert.NoError(t, err)
	assert.Equal(t, morningWorkoutStatistics[:2], fetchedTaskStatistics)
}

// сколько за прошедшую неделю
func TestMockTaskStatisticRepository_FetchAllTaskStatisticsForLastWeek(t *testing.T) {
	morningWorkout, err := CreateMorningWorkoutTask(t)

	taskStatisticRepo := NewMockTaskStatisticRepository()

	morningWorkoutStatistics := CreateMorningWorkoutStatistics(t, morningWorkout, err, taskStatisticRepo)

	//from := time.Date(2023, 4, 5, 0, 0, 0, 0, time.UTC)
	to := time.Now()
	from := to.AddDate(0, 0, -7)
	fetchedTaskStatistics, err := taskStatisticRepo.FetchAllTaskStatisticsForDays(morningWorkout.ID, from, to)
	assert.NoError(t, err)
	assert.Equal(t, morningWorkoutStatistics, fetchedTaskStatistics)
}

func CreateMorningWorkoutTask(t *testing.T) (task.Task, error) {
	taskRepo := repository.NewMockTaskRepository()

	morningWorkout := task.Task{
		ID:       1,
		Name:     "Morning workout",
		TimeGoal: 3 * time.Minute,
		GoalID:   nil,
	}

	err := taskRepo.CreateTask(&morningWorkout)
	assert.NoError(t, err)
	return morningWorkout, err
}

func CreateMorningWorkoutStatistics(t *testing.T, morningWorkout task.Task, err error, taskStatisticRepo *mockTaskStatisticRepository) []task_statistic.TaskStatistic {
	morningWorkoutStatistics := []task_statistic.TaskStatistic{
		{
			ID:        1,
			TaskID:    morningWorkout.ID,
			TimeSpent: 2*time.Minute + 30*time.Second,
			CreatedAt: time.Date(2023, 4, 11, 10, 48, 0, 0, time.UTC),
		},
		{
			ID:        1,
			TaskID:    morningWorkout.ID,
			TimeSpent: 1*time.Minute + 56*time.Second,
			CreatedAt: time.Date(2023, 4, 10, 8, 15, 0, 0, time.UTC),
		},
		{
			ID:        1,
			TaskID:    morningWorkout.ID,
			TimeSpent: 1*time.Minute + 56*time.Second,
			CreatedAt: time.Date(2023, 4, 9, 9, 23, 0, 0, time.UTC),
		},
	}

	for _, morningWorkoutStatistic := range morningWorkoutStatistics {
		err = taskStatisticRepo.CreateTaskStatistic(&morningWorkoutStatistic)
		assert.NoError(t, err)
	}

	return morningWorkoutStatistics
}
