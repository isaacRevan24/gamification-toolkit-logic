package model

const (
	SUCCESS_CODE_STATUS          = "CODE-00"
	INTERNAL_SERVER_ERROR_STATUS = "CODE-01"
	BAD_REQUEST_ERROR_STATUS     = "CODE-02"
)

type SignUpRequest struct {
	ID string `json:"id" binding:"required"`
}

type SignUpResponse struct {
	Status StatusResponse `json:"status"`
}

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

type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
