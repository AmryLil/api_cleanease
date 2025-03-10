package helpers

import (
	"errors"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type validation struct{}

func NewValidationRequest() ValidationInterface {
	return &validation{}
}

func (v validation) ValidateRequest(request any) ([]string, error) {
	var validate = validator.New()

	validate.RegisterValidation("fullname", func(fl validator.FieldLevel) bool {
		regex := regexp.MustCompile("^[a-zA-Z ]+$")
		return regex.MatchString(fl.Field().String())
	})

	if err := validate.Struct(request); err != nil {
		var errMap = []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errMap = append(errMap, strings.ToLower(err.Error()))
		}

		return errMap, errors.New("validation failed")
	}

	return nil, nil
}
