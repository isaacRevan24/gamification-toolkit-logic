package model

import "time"

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

type DeleteHabitResponse struct {
	Status StatusResponse `json:"status"`
}

type CheckHabitHeaders struct {
	UserId  string `header:"user-id"`
	HabitId int    `header:"habit-id"`
	// Time format time.RFC3339
	Date time.Time `header:"date" time_format:"2006-01-02T15:04:05Z07:00"`
}
