package model

type SignUpRequest struct {
	ID string `json:"id" binding:"required"`
}

type SignUpResponse struct {
	Status StatusResponse `json:"status"`
}
