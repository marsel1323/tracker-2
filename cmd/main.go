package main

import (
	"log"
	"time"
	"tracker-2/internals/app/task"
	"tracker-2/internals/model"
)

func main() {
	developmentCategory := model.Category{
		ID:   1,
		Name: "development",
	}
	log.Printf("developmentCategory: %+v", developmentCategory)

	learnLanguagesCategory := model.Category{
		ID:   2,
		Name: "learn languages",
	}
	log.Printf("learnLanguagesCategory: %+v", learnLanguagesCategory)

	healthAndFitnessCategory := model.Category{
		ID:   3,
		Name: "Health and Fitness",
	}
	log.Printf("healthAndFitnessCategory: %+v", healthAndFitnessCategory)

	categories := []model.Category{
		developmentCategory,
		learnLanguagesCategory,
		healthAndFitnessCategory,
	}
	log.Printf("categories: %+v", categories)

	findAJobGoal := model.Goal{
		ID:           1,
		Title:        "Find a job as a software developer with salary of 5000$ per month",
		ParentGoalID: nil,
		CategoryID:   &developmentCategory.ID,
	}
	log.Printf("findAJobGoal: %+v", findAJobGoal)

	learnAlgsAndDSGoal := model.Goal{
		ID:           2,
		Title:        "Learn algorithms and data structures",
		ParentGoalID: &findAJobGoal.ID,
		CategoryID:   &developmentCategory.ID,
	}
	log.Printf("learnAlgsAndDSGoal: %+v", learnAlgsAndDSGoal)

	improveEnglishGoal := model.Goal{
		ID:           3,
		Title:        "Improve English from B1 to C1",
		ParentGoalID: &findAJobGoal.ID,
		CategoryID:   &learnLanguagesCategory.ID,
	}
	log.Printf("improveEnglishGoal: %+v", improveEnglishGoal)

	exerciseEveryMorningGoal := model.Goal{
		ID:           4,
		Title:        "Exercise every morning",
		ParentGoalID: nil,
		CategoryID:   &healthAndFitnessCategory.ID,
	}

	goals := []model.Goal{
		findAJobGoal,
		learnAlgsAndDSGoal,
		improveEnglishGoal,
		exerciseEveryMorningGoal,
	}
	log.Printf("goals: %+v", goals)

	morningWorkout := task.Task{
		ID:       1,
		Name:     "Morning workout",
		TimeGoal: 3 * time.Minute,
		GoalID:   &exerciseEveryMorningGoal.ID,
	}
	log.Printf("morningWorkout: %+v", morningWorkout)

	morningWorkoutStats := model.TaskStatistic{
		ID:        1,
		TaskID:    morningWorkout.ID,
		TimeSpent: 5*time.Minute + 30*time.Second,
		CreatedAt: time.Now(),
	}
	log.Printf("morningWorkoutStats: %+v", morningWorkoutStats)
}

// Testcase:
// есть задача и её статистика за несколько дней
// нужно вывести сколько всего времени потрачено
// нужно посчитать сколько времени потрачено сегодня
