package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/train-cat/api-train/app/models"
)

type user struct{}

// User namespace
var User user

// Persist user in database
func (r user) Persist(i *models.UserInput) (*models.User, error) {
	u, err := i.ToEntity()

	if err != nil {
		return nil, err
	}

	err = db.Save(u).Error

	return u, err
}

// FindByUsername an user
func (r user) FindByUsername(username string) (*models.User, error) {
	u := &models.User{}

	err := db.Where("username = ?", username).First(u).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return u, err
}
