package validations

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationError(err error) *[]ValidationError {
	var validaErrs []ValidationError
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, er := range err.(validator.ValidationErrors) {
			var validationErr ValidationError
			validationErr.Property = er.Field()
			validationErr.Tag = er.Tag()
			validationErr.Value = er.Param()
			validationErr.Message = er.Error()

			validaErrs = append(validaErrs, validationErr)
		}

		return &validaErrs
	}
	return nil

}
