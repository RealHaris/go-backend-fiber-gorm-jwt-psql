package errors

import "github.com/gofiber/fiber/v2"

// NewErrorResponse constructs a generic error response
func NewErrorResponse(message string) fiber.Map {
    return fiber.Map{"error": message}
}

// ErrorResponse DTO for error responses
type ErrorResponse struct {
    Error string `json:"error"`
}
