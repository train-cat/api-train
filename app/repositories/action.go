package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/train-cat/api-train/app/models"
)

type action struct{}

// Action namespace
var Action action

// Persist one action in database
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

// FindOne action by id
func (r action) FindOne(id uint) (*models.Action, error) {
	a := &models.Action{}

	err := db.Where("id = ?", id).First(a).Error

	return a, err
}

// FindByUUID action
func (r action) FindByUUID(uuid string) (*models.Action, error) {
	e := &models.Action{}

	err := db.Where("uuid = ?", uuid).First(e).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return e, err
}
