package utils

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, code int, err error) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
		Err:     err,
	}
}

func RespondWithError(w http.ResponseWriter, appErr *AppError) {
	Logger.Error("Ошибка",
		zap.Error(appErr.Err),
		zap.Int("status", appErr.Code),
		zap.String("message", appErr.Message),
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.Code)
	_ = json.NewEncoder(w).Encode(appErr)
}
