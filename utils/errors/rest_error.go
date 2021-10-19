package errors

import "net/http"

type RestErr struct {
	Message string `json:"message,omitempty"`
	Status    int    `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewBadRequestError(message string ) *RestErr  {
	return &RestErr{
		Message: message,
		Status:    http.StatusBadRequest,
		Error:   "错误请求",
	}
}
func NewNotFoundError(message string ) *RestErr  {
	return &RestErr{
		Message: message,
		Status:    http.StatusNotFound,
		Error:   "未找到",
	}
}