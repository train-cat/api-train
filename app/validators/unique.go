package validators

import (
	"aahframework.org/log.v0"
	"github.com/train-cat/api-train/app/repositories"
	"gopkg.in/go-playground/validator.v9"
)

const tagUnique = "unique"

func uniqueConstraint(f validator.FieldLevel) bool {
	top := f.Top()

	if !top.IsValid() {
		return false
	}

	exist, err := repositories.ValueExist(top.Interface(), f.Param(), f.Field().String())

	if err != nil {
		log.Error(err)
		return false
	}

	return !exist
}
