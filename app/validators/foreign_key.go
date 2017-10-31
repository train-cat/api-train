package validators

import (
	"aahframework.org/log.v0"
	"github.com/train-cat/api-train/app/repositories"
	"gopkg.in/go-playground/validator.v9"
)

const tagForeignKey = "foreign_key"

func ForeignKeyConstraint(f validator.FieldLevel) bool {
	exist, err := repositories.IdExist(f.Param(), f.Field().Int())

	if err != nil {
		log.Error(err)
		return false
	}

	return exist
}