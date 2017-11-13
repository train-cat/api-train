package validators

import (
	"github.com/train-cat/api-train/app/models"
	"gopkg.in/go-playground/validator.v9"
)

const tagActionType = "action_type"

func actionTypeConstraint(f validator.FieldLevel) bool {
	typ := f.Field().String()

	for _, t := range models.AllActionTypes {
		if typ == t {
			return true
		}
	}

	return false
}
