package validators

import "gopkg.in/go-playground/validator.v9"

const tagTrue = "true"

func TrueConstraint(f validator.FieldLevel) bool {
	return f.Field().Bool()
}
