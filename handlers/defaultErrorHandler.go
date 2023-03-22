package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

// DefaultErrorHandler Default error handler
func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errorCode": code,
		})
	}

	// Return from handler
	return nil
}
