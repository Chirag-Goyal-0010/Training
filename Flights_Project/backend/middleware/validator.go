package middleware

import (
	"flights-project/logger"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var validate *validator.Validate

// InitValidator initializes the validator with custom validations
func InitValidator() {
	validate = validator.New()

	// Register custom validations
	registerCustomValidations()
}

// registerCustomValidations registers custom validation rules
func registerCustomValidations() {
	// Register custom validation for flight codes (e.g., "AA123")
	validate.RegisterValidation("flightcode", func(fl validator.FieldLevel) bool {
		code := fl.Field().String()
		if len(code) < 2 || len(code) > 6 {
			return false
		}
		// First two characters should be letters
		if !isLetter(code[0]) || !isLetter(code[1]) {
			return false
		}
		// Rest should be numbers
		for i := 2; i < len(code); i++ {
			if !isNumber(code[i]) {
				return false
			}
		}
		return true
	})

	// Register custom validation for dates (must be future dates)
	validate.RegisterValidation("futuredate", func(fl validator.FieldLevel) bool {
		// Add your date validation logic here
		return true // Placeholder
	})

	// Register custom validation for phone numbers
	validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()
		// Remove any non-digit characters
		phone = strings.Map(func(r rune) rune {
			if r >= '0' && r <= '9' {
				return r
			}
			return -1
		}, phone)
		return len(phone) >= 10 && len(phone) <= 15
	})
}

// isLetter checks if a byte is a letter
func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// isNumber checks if a byte is a number
func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateRequest validates the request body against the provided struct
func ValidateRequest(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind the request body to the model
		if err := c.ShouldBindJSON(model); err != nil {
			c.Error(BadRequestError("Invalid request body", err))
			return
		}

		// Validate the model
		if err := validate.Struct(model); err != nil {
			validationErrors := make([]ValidationError, 0)

			// Convert validation errors to a more user-friendly format
			for _, err := range err.(validator.ValidationErrors) {
				field := err.Field()
				tag := err.Tag()
				param := err.Param()

				// Get the field's JSON tag if available
				if t, ok := reflect.TypeOf(model).FieldByName(field); ok {
					if jsonTag := t.Tag.Get("json"); jsonTag != "" {
						field = strings.Split(jsonTag, ",")[0]
					}
				}

				// Create user-friendly error message
				message := getValidationMessage(field, tag, param)
				validationErrors = append(validationErrors, ValidationError{
					Field:   field,
					Message: message,
				})
			}

			// Log validation errors
			logger.Error("Validation failed",
				zap.Any("errors", validationErrors),
				zap.String("path", c.Request.URL.Path),
			)

			// Return validation errors
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errors":  validationErrors,
			})
			c.Abort()
			return
		}

		// Store the validated model in the context
		c.Set("validatedModel", model)
		c.Next()
	}
}

// getValidationMessage returns a user-friendly validation message
func getValidationMessage(field, tag, param string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return "Invalid email format"
	case "min":
		return field + " must be at least " + param + " characters"
	case "max":
		return field + " must not exceed " + param + " characters"
	case "flightcode":
		return "Invalid flight code format (e.g., AA123)"
	case "futuredate":
		return "Date must be in the future"
	case "phone":
		return "Invalid phone number format"
	default:
		return field + " failed validation: " + tag
	}
}

// GetValidatedModel retrieves the validated model from the context
func GetValidatedModel(c *gin.Context, model interface{}) bool {
	if val, exists := c.Get("validatedModel"); exists {
		reflect.ValueOf(model).Elem().Set(reflect.ValueOf(val).Elem())
		return true
	}
	return false
}
