package structures

import "github.com/go-playground/validator/v10"

//Validate validates item request body
func (s *ItemValidator) Validate(i interface{}) error {
	return s.validator.Struct(i)
}

func NewSampValidator() *ItemValidator {
	var v = validator.New()
	var sv = &ItemValidator{validator: v}
	return sv
}
