package repositories

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/jinzhu/gorm"
)

type action struct{}

var Action action

func (r action) Persist(i *models.ActionInput) (*models.Action, error) {
	a := i.ToEntity()

	// if already exist return it
	if exist, err := ValueExist(a, "uuid", a.UUID); exist {
		if err != nil {
			return nil, err
		}

		a, err = r.FindByUUID(a.UUID)

		if err != nil {
			return nil, err
		}

		if a != nil {
			return a, err
		}
	}

	err := db.Save(a).Error

	return a, err
}

func (_ action) FindByUUID(uuid string) (*models.Action, error) {
	e := &models.Action{}

	err := db.Where("uuid = ?", uuid).First(e).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return e, err
}
