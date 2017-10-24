package validators

import "gopkg.in/go-playground/validator.v9"

var validate = validator.New()

func init() {
	validate.RegisterValidation(tagUnique, UniqueConstraint)
}

func Struct(s interface{}) error {
	return validate.Struct(s)
}
