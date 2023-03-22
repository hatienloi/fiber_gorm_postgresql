package validations

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/interfaces"
)

type ValidatorResponseInterface struct {
	ErrorDetail []interfaces.ErrorResponse `json:"errorDetail"`
	ErrorCode   int                        `json:"errorCode"`
}

func ValidatorResponse(error []interfaces.ErrorResponse) (ValidatorResponseInterface, error) {
	var resp ValidatorResponseInterface
	resp.ErrorCode = fiber.StatusBadRequest
	resp.ErrorDetail = error
	if error != nil {
		return resp, errors.New("validations error")
	}
	return resp, nil
}
