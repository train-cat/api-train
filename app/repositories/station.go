package repositories

import (
	"strconv"

	"aahframework.org/log.v0"
	"github.com/jinzhu/gorm"
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/models"
)

type station struct{}

var Station station

func (_ station) Persist(i *models.StationInput) (*models.Station, error) {
	s := i.ToEntity()

	err := db.Save(s).Error

	return s, err
}

func (_ station) FindAll(f filters.Filter) (*models.Collection, error) {
	stations := []*models.Station{}

	db := db.Model(models.Station{}).Where("is_realtime = 1")

	return NewCollection(f, db, &stations)
}

func (_ station) FindOne(stationID uint) (*models.Station, error) {
	s := &models.Station{}

	err := db.Where("id = ?", stationID).First(s).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return s, err
}

func (_ station) IsExist(stationID int) bool {
	exist, err := ValueExist(&models.Station{}, "id", strconv.Itoa(stationID))

	if err != nil {
		log.Error(err)
	}

	return exist
}
