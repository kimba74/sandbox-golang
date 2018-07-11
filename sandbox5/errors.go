package main

import (
	"fmt"
	"time"
)

// AppError is an error structure for sandbox5
type AppError struct {
	message   string
	code      int
	timeStamp time.Time
}

// NewError creates a new error object
func NewError(msg string, code int) *AppError {
	return &AppError{
		message:   msg,
		code:      code,
		timeStamp: time.Now(),
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf(e.message)
}
