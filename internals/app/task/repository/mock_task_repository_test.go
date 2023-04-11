package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"tracker-2/internals/app/task"
)

func TestCreateTask(t *testing.T) {
	taskRepo := NewMockTaskRepository()

	morningWorkout := task.Task{
		ID:       1,
		Name:     "Morning workout",
		TimeGoal: 3 * time.Minute,
		GoalID:   nil,
	}

	err := taskRepo.CreateTask(&morningWorkout)
	if err != nil {
		t.Errorf("Error while creating task: %v", err)
	}
	assert.NoError(t, err)
}

func TestFetchByID(t *testing.T) {
	taskRepo := NewMockTaskRepository()

	morningWorkout := task.Task{
		ID:       1,
		Name:     "Morning workout",
		TimeGoal: 3 * time.Minute,
		GoalID:   nil,
	}

	err := taskRepo.CreateTask(&morningWorkout)
	if err != nil {
		t.Errorf("Error while creating task: %v", err)
	}
	assert.NoError(t, err)

	fetchedTask, err := taskRepo.FetchTaskByID(morningWorkout.ID)
	if err != nil {
		t.Errorf("Error while fetching task: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, morningWorkout, *fetchedTask)
}

func TestFetchAll(t *testing.T) {
	taskRepo := NewMockTaskRepository()

	morningWorkout := task.Task{
		ID:       1,
		Name:     "Morning workout",
		TimeGoal: 3 * time.Minute,
		GoalID:   nil,
	}

	err := taskRepo.CreateTask(&morningWorkout)
	if err != nil {
		t.Errorf("Error while creating task: %v", err)
	}
	assert.NoError(t, err)

	eveningWalk := task.Task{
		ID:       2,
		Name:     "Evening walk",
		TimeGoal: 30 * time.Minute,
		GoalID:   nil,
	}

	err = taskRepo.CreateTask(&eveningWalk)
	if err != nil {
		t.Errorf("Error while creating task: %v", err)
	}
	assert.NoError(t, err)

	fetchedTasks, err := taskRepo.FetchTasks()
	if err != nil {
		t.Errorf("Error while fetching tasks: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, 2, len(fetchedTasks))
	assert.Equal(t, morningWorkout, fetchedTasks[0])
	assert.Equal(t, eveningWalk, fetchedTasks[1])
}
