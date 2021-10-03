package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
)

func (*habitController) AddNewHabit(reques model.AddNewHabitRequest) (int, error) {
	const functionName string = "AddNewHabitController"
	Logs.LogDebug("Start " + functionName)

	conditionError := validateHabitCondition(reques.Condition)

	if conditionError != nil {
		Logs.LogError(conditionError)
		Logs.LogDebug("End " + functionName)
		return 0, conditionError
	}

	habitId, err := habitRepository.AddNewHabitRepository(reques.UserId, reques.Name, reques.Description, reques.Condition, reques.Repetition)

	if err != nil {
		Logs.LogDebug("End " + functionName)
		return 0, err
	}

	Logs.LogInfo("New habit created")
	Logs.LogDebug("End " + functionName)

	return habitId, nil
}

func validateHabitCondition(condition string) error {

	for _, i := range conditionTypes {
		if condition == i {
			return nil
		}
	}

	return ErrorInvalidHabitCondition
}
