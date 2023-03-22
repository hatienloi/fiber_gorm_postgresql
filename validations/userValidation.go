package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/interfaces"
)

var validate = validator.New()

func ValidateStruct(user interfaces.User) (ValidatorResponseInterface, error) {
	var errors []interfaces.ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element interfaces.ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Message = err.Error()
			errors = append(errors, element)
		}
	}
	return ValidatorResponse(errors)
}
