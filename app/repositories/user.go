package repositories

import (
	"github.com/train-cat/api-train/app/models"
	"github.com/jinzhu/gorm"
)

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

func (_ user) FindByUsername(username string) (*models.User, error) {
	u := &models.User{}

	err := db.Where("username = ?", username).First(u).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return u, err
}
