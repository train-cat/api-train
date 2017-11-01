package validators

import "gopkg.in/go-playground/validator.v9"

var validate = validator.New()

func init() {
	validate.RegisterValidation(tagUnique, uniqueConstraint)
	validate.RegisterValidation(tagForeignKey, foreignKeyConstraint)
	validate.RegisterValidation(tagTrue, trueConstraint)
	validate.RegisterValidation(tagActionType, actionTypeConstraint)
}

// Struct valid the structure
func Struct(s interface{}) error {
	return validate.Struct(s)
}
