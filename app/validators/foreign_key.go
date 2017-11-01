package validators

import (
	"aahframework.org/log.v0"
	"github.com/train-cat/api-train/app/repositories"
	"gopkg.in/go-playground/validator.v9"
)

const tagForeignKey = "foreign_key"

func foreignKeyConstraint(f validator.FieldLevel) bool {
	exist, err := repositories.IDExist(f.Param(), f.Field().Uint())

	if err != nil {
		log.Error(err)
		return false
	}

	return exist
}
