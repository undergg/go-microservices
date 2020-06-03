package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

func (p *Product) Validate() error {
	validator := validator.New()

	validator.RegisterValidation("sku", validateSKU)

	// Just takes a look at the struct tags and makes sure everything
	// is according to the constraints.
	return validator.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)

	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}
