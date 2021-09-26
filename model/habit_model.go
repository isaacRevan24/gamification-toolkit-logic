package model

type AddNewHabitRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Condition   string `json:"condition" binding:"required"`
	Repetition  string `json:"repetition" binding:"required"`
}

type AddNewHabitResponse struct {
	Status  StatusResponse `json:"status"`
	HabitId int            `json:"habit-id"`
}
