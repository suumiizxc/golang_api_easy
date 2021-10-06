package helper

import "strings"

type Response struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors"`
	Data       interface{} `json:"data"`
	TotalCount interface{} `json:"total_count"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:     status,
		Message:    message,
		Errors:     nil,
		Data:       data,
		TotalCount: nil,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:     false,
		Message:    message,
		Errors:     splittedError,
		Data:       data,
		TotalCount: nil,
	}
	return res
}
