package structures

import (
	"github.com/go-playground/validator/v10"
)

//Item is the struct that contains the json fields. Simple Name, Description.
type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//echo validate, use custom interface
type ItemValidator struct {
	validator *validator.Validate
}
