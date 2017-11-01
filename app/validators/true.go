package validators

import "gopkg.in/go-playground/validator.v9"

const tagTrue = "true"

func trueConstraint(f validator.FieldLevel) bool {
	return f.Field().Bool()
}
