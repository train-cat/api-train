package validators

import (
	"gopkg.in/go-playground/validator.v9"
	"github.com/train-cat/api-train/app/models"
)

const tagActionType = "action_type"

func ActionTypeConstraint(f validator.FieldLevel) bool {
	typ := f.Field().String()

	for _, t := range models.AllActionTypes {
		if typ == t {
			return true
		}
	}

	return false
}
