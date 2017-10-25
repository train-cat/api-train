package repositories

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/jinzhu/gorm"
)

type train struct{}

var Train train

func (_ train) Persist(i *models.TrainInput) (*models.Train, error) {
	t := i.ToEntity()

	err := db.Save(t).Error

	return t, err
}

func (r train) FindOneByCode(code string) (*models.Train, error) {
	train := &models.Train{}

	err := db.
	Preload("Terminus").
		Where("code = ?", code).First(train).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return train, nil
}
