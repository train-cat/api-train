package repositories

import (
	"strconv"

	"aahframework.org/log.v0"
	"github.com/jinzhu/gorm"
	"github.com/train-cat/api-train/app/filters"
	"github.com/train-cat/api-train/app/models"
)

type station struct{}

// Station namespace
var Station station

// Persist station in database
func (r station) Persist(i *models.StationInput) (*models.Station, error) {
	s := i.ToEntity()

	err := db.Save(s).Error

	return s, err
}

// FindAll stations
func (r station) FindAll(f filters.Filter) (*models.Collection, error) {
	stations := []*models.Station{}

	db := db.Model(models.Station{}).Where("is_enable = 1")

	return NewCollection(f, db, &stations)
}

// FindOne station by id
func (r station) FindOne(stationID uint) (*models.Station, error) {
	s := &models.Station{}

	err := db.Where("id = ?", stationID).First(s).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return s, err
}

// IsExist return true if stationID exist
func (r station) IsExist(stationID int) bool {
	exist, err := ValueExist(&models.Station{}, "id", strconv.Itoa(stationID))

	if err != nil {
		log.Error(err)
	}

	return exist
}

// FindTerminus return terminus of the trips
func (r station) FindTerminus(tripID uint) (*models.Station, error) {
	s := &models.Station{}

	err := db.
		Joins("LEFT JOIN stop_time ON stop_time.station_id = station.id").
		Where("stop_time.trip_id = ?", tripID).
		Order("stop_time.schedule DESC").
		First(s).Error

	return s, err
}
