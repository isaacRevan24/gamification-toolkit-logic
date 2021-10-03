package controller

import (
	"github.com/isaacRevan24/gamification-toolkit-logic/model"
)

func (*userController) SignUp(request model.SignUpRequest) error {
	const functionName string = "SignUpController"
	Logs.LogDebug("Start " + functionName)

	err := userRepository.SignUpRepository(request.ID)

	if err != nil {
		Logs.LogDebug("End " + functionName)
		return err
	}

	Logs.LogInfo("Successfully saved user.")
	Logs.LogDebug("End " + functionName)

	return nil
}
