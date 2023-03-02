package fiberw

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Validator Error Responses
type ValidationErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

// Standard API error format
type RequestError struct {
	StatusCode int
	Err        error
	Message    string
}

// Error Impl
func (re *RequestError) Error() string {
	return fmt.Sprintf("Request Errored Out %s", re.Err)
}

// Helper to create a new Request Error
func NewRequestError(statusCode int, message string, err error) *RequestError {
	return &RequestError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

// Assign the request body to the given interface and then check for validity
func ValidateBody(data interface{}, c *fiber.Ctx) (bool, []*ValidationErrorResponse) {
	if err := c.BodyParser(&data); err != nil {
		errData := make([]*ValidationErrorResponse, 1)
		error := ValidationErrorResponse{
			FailedField: "nil",
			Tag:         "Parse Error",
			Value:       "Check Logs",
		}
		log.Println(err)
		errData = append(errData, &error)
		return true, errData
	}

	var validate = validator.New()
	var errors []*ValidationErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	if len(errors) > 0 {
		return true, errors
	}
	return false, nil
}

// Handle RequestErrors
func handleError(err error, ctx *fiber.Ctx) error {
	if reqError, ok := err.(*RequestError); ok {
		return ctx.Status(reqError.StatusCode).JSON(fiber.Map{
			"status_code": reqError.StatusCode,
			"error":       reqError.Err.Error(),
			"message":     reqError.Message,
		})
	} else {
		panic(err)
	}
}
