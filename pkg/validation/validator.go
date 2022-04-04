package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator defintion
var Validator *validator.Validate

// Initialize validator
func init() {
	Validator = validator.New()
}
