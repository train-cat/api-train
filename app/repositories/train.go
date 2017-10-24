package repositories

import "github.com/train-sh/api-train/app/models"

type train struct{}

var Train train

func (_ train) Persist(i *models.TrainInput) (*models.Train, error) {
	t := i.ToEntity()

	err := db.Save(t).Error

	return t, err
}
