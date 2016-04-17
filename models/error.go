package models

import "fmt"

type Error interface {
	error
	StatusCode() int
	Message() string
}

type errorResponse struct {
	statusCode int
	message    string
}

func NewError(code int, message string) Error {
	return &errorResponse{code, message}
}

func (e *errorResponse) Error() string {
	return fmt.Sprintf("code:%d, message:%s", e.statusCode, e.message)
}
func (e *errorResponse) StatusCode() int {
	return e.statusCode
}
func (e *errorResponse) Message() string {
	return e.message
}
