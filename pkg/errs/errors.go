package errs

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	error   `json:"error,omitempty"`
	Code    int    `json:"status_code,omitempty"`
	Message string `json:"message"`
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewInvalidParameterError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewInvalidBodyError(message string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewInvalidRequestMethodError(message string) *AppError {
	return &AppError{
		Code:    http.StatusMethodNotAllowed,
		Message: message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func (e *AppError) ToJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		return
	}
}
