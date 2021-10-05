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

func (*habitController) DeleteHabit(userId string, habitId string) (bool, error) {
	const functionName string = "DeleteHabit"
	Logs.LogDebug("Start " + functionName)

	response, err := habitRepository.DeleteHabitRepository(userId, habitId)

	if err != nil {
		Logs.LogDebug("End " + functionName)
		return false, err
	}

	Logs.LogDebug("End " + functionName)
	return response, nil
}

func validateHabitCondition(condition string) error {

	conditionTypes := [3]string{"D", "W", "Y"}

	for _, i := range conditionTypes {
		if condition == i {
			return nil
		}
	}

	return ErrorInvalidHabitCondition
}
