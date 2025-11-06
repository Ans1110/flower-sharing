package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationError(c *gin.Context) {
	// Check if validation errors exist in context
	validationErrorsRaw, exists := c.Get("validationErrors")
	if !exists {
		c.Next()
		return
	}

	// Get the validation errors
	validationErrors, ok := validationErrorsRaw.(validator.ValidationErrors)
	if !ok || len(validationErrors) == 0 {
		c.Next()
		return
	}

	// Format errors
	errorsMap := make(map[string]interface{})
	for _, err := range validationErrors {
		field := err.Field()
		if _, exists := errorsMap[field]; !exists {
			errorsMap[field] = map[string]interface{}{
				"message": getValidationErrorMessage(err),
				"value":   err.Value(),
			}
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"code":   "ValidationError",
		"errors": errorsMap,
	})
	c.Abort()
}

// ExtractValidationErrors extracts validation errors from a binding error and stores them in context
// This should be called after ShouldBindJSON if it returns an error
// Returns true if validation errors were found and stored, false otherwise
func ExtractValidationErrors(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	// Try to extract validator.ValidationErrors from the binding error
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		c.Set("validationErrors", validationErrors)
		return true
	}

	// Direct type assertion as fallback
	if ve, ok := err.(validator.ValidationErrors); ok {
		c.Set("validationErrors", ve)
		return true
	}

	return false
}

// getValidationErrorMessage returns a human-readable error message
func getValidationErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "len":
		return "Length must be exactly " + err.Param()
	default:
		return "Validation failed for " + err.Tag()
	}
}
