package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/train-cat/api-train/app/models"
)

type trip struct{}

// Trip namespace
var Trip trip

// Find trip
func (r trip) Find(id int) (*models.Trip, error) {
	trip := &models.Trip{}

	err := db.
		Preload("Calendar").
		Where("id = ?", id).First(trip).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return trip, nil
}
