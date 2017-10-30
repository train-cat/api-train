package repositories

import (
	"aahframework.org/log.v0"
	"github.com/train-cat/api-train/app/models"
	"github.com/jinzhu/gorm"
)

type train struct{}

var Train train

func (_ train) Persist(i *models.TrainInput) (*models.Train, error) {
	t := i.ToEntity()

	err := db.Save(t).Error

	if err == nil && t.TerminusID != nil && *t.TerminusID > 0 {
		t.Terminus, err = Station.FindOne(*t.TerminusID)
	}

	return t, err
}

func (_ train) FindOneByCode(code string) (*models.Train, error) {
	train := &models.Train{}

	err := db.
	Preload("Terminus").
		Where("code = ?", code).First(train).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return train, nil
}

func (_ train) IsExist(code string) bool {
	exist, err := ValueExist(&models.Train{}, "code", code)

	if err != nil {
		log.Error(err)
	}

	return exist
}
