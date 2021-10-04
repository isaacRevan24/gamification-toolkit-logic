package model

type AddNewHabitRequest struct {
	UserId      string `json:"user-id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Condition   string `json:"condition" binding:"required"`
	Repetition  int    `json:"repetition" binding:"required"`
}

type AddNewHabitResponse struct {
	Status  StatusResponse `json:"status"`
	HabitId int            `json:"habit-id"`
}

type DeleteHabitRequest struct {
	UserId  string `json:"user-id" binding:"required"`
	HabitId int    `json:"habit-id" binding:"required"`
}

type DeleteHabitResponse struct {
	Status StatusResponse `json:"status"`
}
