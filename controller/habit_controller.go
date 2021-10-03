package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
)

func (*habitController) AddNewHabit(reques model.AddNewHabitRequest) (int, error) {
	const functionName string = "AddNewHabitController"
	Logs.LogDebug("Start " + functionName)

	habitId, err := habitRepository.AddNewHabitRepository(reques.UserId, reques.Name, reques.Description, reques.Condition, reques.Repetition)

	if err != nil {
		Logs.LogDebug("End " + functionName)
		return 0, err
	}

	Logs.LogInfo("New habit created")
	Logs.LogDebug("End " + functionName)

	return habitId, nil
}
