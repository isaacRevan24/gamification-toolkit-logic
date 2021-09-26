package model

const (
	SUCCESS_CODE_STATUS          = "CODE-00"
	INTERNAL_SERVER_ERROR_STATUS = "CODE-01"
	BAD_REQUEST_ERROR_STATUS     = "CODE-02"
)

type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
