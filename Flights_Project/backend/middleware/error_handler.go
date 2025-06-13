package middleware

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AppError represents an application error
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

// Error implements the error interface for AppError
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// ErrorResponse represents the structure of error responses
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
}

// ErrorHandler middleware for handling errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *AppError

			// Convert error to AppError if it's not already
			if !errors.As(err, &appErr) {
				appErr = convertToAppError(err)
			}

			// Send error response
			c.JSON(appErr.Code, ErrorResponse{
				Success: false,
				Error:   appErr.Message,
				Code:    appErr.Code,
			})
		}
	}
}

// convertToAppError converts various error types to AppError
func convertToAppError(err error) *AppError {
	// Handle database errors
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &AppError{
			Code:    http.StatusNotFound,
			Message: "Resource not found",
			Err:     err,
		}
	}

	if errors.Is(err, sql.ErrNoRows) {
		return &AppError{
			Code:    http.StatusNotFound,
			Message: "Resource not found",
			Err:     err,
		}
	}

	// Handle validation errors
	if strings.Contains(err.Error(), "validation failed") {
		return &AppError{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Err:     err,
		}
	}

	// Handle authentication errors
	if strings.Contains(err.Error(), "unauthorized") {
		return &AppError{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized access",
			Err:     err,
		}
	}

	// Handle forbidden errors
	if strings.Contains(err.Error(), "forbidden") {
		return &AppError{
			Code:    http.StatusForbidden,
			Message: "Access forbidden",
			Err:     err,
		}
	}

	// Default error
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Err:     err,
	}
}

// NewAppError creates a new AppError
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Common error constructors
func NotFoundError(message string, err error) *AppError {
	return NewAppError(http.StatusNotFound, message, err)
}

func BadRequestError(message string, err error) *AppError {
	return NewAppError(http.StatusBadRequest, message, err)
}

func UnauthorizedError(message string, err error) *AppError {
	return NewAppError(http.StatusUnauthorized, message, err)
}

func ForbiddenError(message string, err error) *AppError {
	return NewAppError(http.StatusForbidden, message, err)
}

func InternalServerError(message string, err error) *AppError {
	return NewAppError(http.StatusInternalServerError, message, err)
}
