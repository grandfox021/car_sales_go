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

func Get_validation_errors(err error) *[]ValidationError {

	var validationerros []ValidationError
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _,err := range err.(validator.ValidationErrors){
			var el ValidationError
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationerros = append(validationerros, el)
		}
		return &validationerros
	}
	return  nil
}