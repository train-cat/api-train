package repositories

import "github.com/train-cat/api-train/app/models"

type user struct {}

var User user

func (_ user) Persist(i *models.UserInput) (*models.User, error) {
	u, err := i.ToEntity()

	if err != nil {
		return nil, err
	}

	err = db.Save(u).Error

	return u, err
}
